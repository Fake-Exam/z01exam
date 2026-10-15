package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

const (
	javaCatalogURL        = "https://zone01normandie.org/api/object/rouen"
	officialJavaTestImage = "ghcr.io/01-edu/test-java:latest"
)

// LoadJavaExercises deliberately exposes the three official Java checkpoints,
// rather than the whole curriculum. Zone01 currently publishes no Java final
// exam object in this catalogue.
func LoadJavaExercises() ([]Exercise, error) {
	ctx, cancel := context.WithTimeout(context.Background(), runTimeout())
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, javaCatalogURL, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("catalogue Java: HTTP %d", res.StatusCode)
	}
	var root zoneRustNode
	if err := json.NewDecoder(res.Body).Decode(&root); err != nil {
		return nil, err
	}
	targets := map[string]string{
		"/rouen/div-01/piscine-java/checkpoint-01-java": "Checkpoint 01",
		"/rouen/div-01/piscine-java/checkpoint-02-java": "Checkpoint 02",
		"/rouen/div-01/piscine-java/checkpoint-03-java": "Checkpoint 03",
	}
	var list []Exercise
	var visit func(zoneRustNode)
	visit = func(node zoneRustNode) {
		if label, ok := targets[node.Path]; ok {
			for _, child := range node.Children {
				if child.Type != "exercise" || child.Key == "" {
					continue
				}
				runner := strings.TrimSpace(child.Name)
				if runner == "" {
					runner = javaClassName(child.Key)
				}
				image := officialJavaTestImage
				for _, validation := range child.Attrs.Validations {
					if validation.Type == "tester" && validation.TestImage != "" {
						image = validation.TestImage
						break
					}
				}
				file := runner + ".java"
				list = append(list, Exercise{ID: "java-" + child.Key, Title: label + " · " + runner, Runner: runner, Template: "public class " + runner + " {\n    // Écris ta solution ici.\n}\n", Statement: javaStatementPlaceholder(child.Attrs.Subject), ExpectedFiles: []string{file}, Image: image, SubjectURL: child.Attrs.Subject})
			}
			return
		}
		for _, child := range node.Children {
			visit(child)
		}
	}
	visit(root)
	sort.Slice(list, func(i, j int) bool { return list[i].Title < list[j].Title })
	hydrateJavaStatements(ctx, list)
	return list, nil
}

func javaClassName(key string) string {
	var out strings.Builder
	for _, part := range strings.FieldsFunc(key, func(r rune) bool { return r == '-' || r == '_' }) {
		if part == "" {
			continue
		}
		out.WriteString(strings.ToUpper(part[:1]))
		out.WriteString(part[1:])
	}
	return out.String()
}

func javaStatementPlaceholder(subject string) string {
	if subject == "" {
		return "## Énoncé Java\n\nÉnoncé officiel indisponible."
	}
	return "## Énoncé Java\n\nChargement de l’énoncé officiel Zone01 : `" + subject + "`."
}

func hydrateJavaStatements(ctx context.Context, exercises []Exercise) {
	jobs := make(chan int)
	var wg sync.WaitGroup
	for range 8 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				if statement, err := fetchZoneStatement(ctx, exercises[i].SubjectURL); err == nil && statement != "" {
					exercises[i].Statement = statement
				}
			}
		}()
	}
	for i := range exercises {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}

func RunJavaExercise(ex Exercise, code string) (bool, string) {
	dir, err := os.MkdirTemp(rootDir, "javaexam-*")
	if err != nil {
		return false, "Erreur système: " + err.Error()
	}
	defer os.RemoveAll(dir)
	_ = os.Chmod(dir, 0o777)
	if err := prepareJavaStudentDir(dir, ex, code); err != nil {
		return false, "Erreur de préparation Java: " + err.Error()
	}
	ctx, cancel := context.WithTimeout(context.Background(), runTimeout())
	defer cancel()
	image := ex.Image
	if image == "" {
		image = officialJavaTestImage
	}
	args := append([]string{"run", "--rm", "--platform", "linux/amd64"}, baseDockerRunArgs("")...)
	args = withOptionalCPULimit(args)
	args = append(args, "-v", dir+":/app/student:rw", "-e", "EXERCISE="+ex.Runner, image)
	cmd := exec.CommandContext(ctx, "docker", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err = cmd.Run()
	out := strings.TrimSpace(stdout.String() + "\n" + stderr.String())
	if ctx.Err() == context.DeadlineExceeded {
		return false, "Temps limite dépassé."
	}
	if err != nil {
		if out == "" {
			out = "Échec moulinette Java officielle: " + err.Error()
		}
		return false, out
	}
	if out == "" {
		out = "OK: tous les tests Java sont validés."
	}
	return true, out
}

func prepareJavaStudentDir(dir string, ex Exercise, code string) error {
	project := filepath.Join(dir, ex.Runner)
	if err := os.MkdirAll(project, 0o755); err != nil {
		return err
	}
	if !strings.HasSuffix(code, "\n") {
		code += "\n"
	}
	return os.WriteFile(filepath.Join(project, ex.Runner+".java"), []byte(code), 0o644)
}
