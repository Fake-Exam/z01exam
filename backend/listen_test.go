package backend

import (
	"bytes"
	"net"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"
)

// Exercise the real server entrypoint: staging must honor LISTEN_ADDR even
// when PORT is unsuitable, without changing the production PORT contract.
func TestRunHonorsLoopbackListenAddress(t *testing.T) {
	if os.Getenv("Z02_TEST_LISTENER_CHILD") == "1" {
		if err := Run("web"); err != nil {
			os.Stderr.WriteString(err.Error())
			os.Exit(1)
		}
		return
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	address := listener.Addr().String()
	listener.Close()
	t.Setenv("Z02_TEST_LISTENER_CHILD", "1")
	t.Setenv("LISTEN_ADDR", address)
	t.Setenv("PORT", "intentionally-invalid-port")
	t.Setenv("CATALOG_SERVICE_URL", "http://127.0.0.1:1")
	t.Setenv("RUNNER_SERVICE_URL", "http://127.0.0.1:1")
	cmd := exec.Command(os.Args[0], "-test.run=^TestRunHonorsLoopbackListenAddress$")
	var output bytes.Buffer
	cmd.Stdout, cmd.Stderr = &output, &output
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	finished := false
	t.Cleanup(func() {
		if !finished {
			cmd.Process.Kill()
			<-done
		}
	})
	client := &http.Client{Timeout: 100 * time.Millisecond}
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		select {
		case err := <-done:
			finished = true
			t.Fatalf("server exited before listening on %s: %v\n%s", address, err, output.String())
		default:
		}
		response, err := client.Get("http://" + address + "/")
		if err == nil {
			response.Body.Close()
			return
		}
		time.Sleep(20 * time.Millisecond)
	}
	t.Fatalf("server did not listen on explicit loopback address %s", address)
}
