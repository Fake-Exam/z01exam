package backend

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServiceModeRequiresAnExplicitComponent(t *testing.T) {
	if _, err := serviceMode(""); err == nil {
		t.Fatal("serviceMode accepted an empty component")
	}
}

func TestServiceModeAcceptsIndependentServices(t *testing.T) {
	for _, mode := range []string{serviceModeCatalog, serviceModeRunner, serviceModeWeb} {
		got, err := serviceMode(mode)
		if err != nil || got != mode {
			t.Fatalf("serviceMode(%q) = %q, %v", mode, got, err)
		}
	}
}

func TestServiceModeRejectsUnknownValue(t *testing.T) {
	if _, err := serviceMode("worker"); err == nil {
		t.Fatal("serviceMode accepted an unknown mode")
	}
}

func TestCatalogAndRunnerExposeOnlyTheirContracts(t *testing.T) {
	catalog, err := newServiceMux(serviceModeCatalog)
	if err != nil {
		t.Fatalf("catalog mux: %v", err)
	}
	if got := httptest.NewRecorder(); func() int {
		catalog.ServeHTTP(got, httptest.NewRequest(http.MethodGet, "/api/run", nil))
		return got.Code
	}() != http.StatusNotFound {
		t.Fatalf("catalog exposes runner route: status %d", got.Code)
	}

	runner, err := newServiceMux(serviceModeRunner)
	if err != nil {
		t.Fatalf("runner mux: %v", err)
	}
	res := httptest.NewRecorder()
	runner.ServeHTTP(res, httptest.NewRequest(http.MethodGet, "/api/exercises", nil))
	if res.Code != http.StatusNotFound {
		t.Fatalf("runner exposes catalog route: status %d", res.Code)
	}
}

func TestMonolithKeepsBothPublicAPIContracts(t *testing.T) {
	monolith, err := newServiceMux(serviceModeMonolith)
	if err != nil {
		t.Fatalf("monolith mux: %v", err)
	}
	catalogResponse := httptest.NewRecorder()
	monolith.ServeHTTP(catalogResponse, httptest.NewRequest(http.MethodGet, "/api/exercises?track=go", nil))
	if catalogResponse.Code != http.StatusOK {
		t.Fatalf("monolith catalog status = %d, want %d", catalogResponse.Code, http.StatusOK)
	}
	runnerResponse := httptest.NewRecorder()
	monolith.ServeHTTP(runnerResponse, httptest.NewRequest(http.MethodGet, "/api/run", nil))
	if runnerResponse.Code != http.StatusMethodNotAllowed {
		t.Fatalf("monolith runner method status = %d, want %d", runnerResponse.Code, http.StatusMethodNotAllowed)
	}
}

func TestWebRequiresBothServiceURLs(t *testing.T) {
	t.Setenv("CATALOG_SERVICE_URL", "")
	t.Setenv("RUNNER_SERVICE_URL", "")
	if _, err := newServiceMux(serviceModeWeb); err == nil {
		t.Fatal("web mux accepted missing service URLs")
	}
}

func TestWebForwardsEachPublicAPIToItsService(t *testing.T) {
	catalog := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/exercises" {
			http.Error(w, "wrong catalog path", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer catalog.Close()
	runner := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/run" {
			http.Error(w, "wrong runner path", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	}))
	defer runner.Close()
	t.Setenv("CATALOG_SERVICE_URL", catalog.URL)
	t.Setenv("RUNNER_SERVICE_URL", runner.URL)

	web, err := newServiceMux(serviceModeWeb)
	if err != nil {
		t.Fatalf("web mux: %v", err)
	}
	for _, tc := range []struct {
		path string
		want int
	}{
		{path: "/api/exercises?track=js", want: http.StatusNoContent},
		{path: "/api/run", want: http.StatusAccepted},
	} {
		res := httptest.NewRecorder()
		web.ServeHTTP(res, httptest.NewRequest(http.MethodGet, tc.path, nil))
		if res.Code != tc.want {
			t.Fatalf("%s: status = %d, want %d", tc.path, res.Code, tc.want)
		}
	}
}

func TestRunnerCatalogPreservesPrivateImageMetadata(t *testing.T) {
	input := []Exercise{{ID: "rust-checkpoint", Image: "ghcr.io/01-edu/test-rust:sha"}}
	encoded, err := json.Marshal(RunnerCatalog{Rust: runnerExercises(input)})
	if err != nil {
		t.Fatalf("marshal RunnerCatalog: %v", err)
	}
	if !strings.Contains(string(encoded), `"image":"ghcr.io/01-edu/test-rust:sha"`) {
		t.Fatalf("private image metadata is missing from RunnerCatalog: %s", encoded)
	}
	decoded := catalogExercises(runnerExercises(input))
	if got := decoded[0].Image; got != input[0].Image {
		t.Fatalf("image after snapshot round-trip = %q, want %q", got, input[0].Image)
	}
}

func TestRunnerLoadsCatalogSnapshotFromCatalogService(t *testing.T) {
	snapshot := RunnerCatalog{Go: []RunnerExercise{{Exercise: Exercise{ID: "onlya", Runner: "onlya"}}}}
	catalog := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/internal/runner-catalog" {
			http.Error(w, "wrong catalog path", http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(snapshot)
	}))
	defer catalog.Close()
	t.Setenv("CATALOG_SERVICE_URL", catalog.URL)
	previousGo, previousJS := exercises, jsExercises
	previousRust, previousJava := rustExercises, javaExercises
	t.Cleanup(func() {
		exercises, jsExercises = previousGo, previousJS
		rustExercises, javaExercises = previousRust, previousJava
	})

	if err := loadRunnerCatalog(); err != nil {
		t.Fatalf("loadRunnerCatalog: %v", err)
	}
	if ex := FindExercise(exercises, "onlya"); ex == nil || ex.Runner != "onlya" {
		t.Fatalf("runner catalog not loaded: %#v", exercises)
	}
}
