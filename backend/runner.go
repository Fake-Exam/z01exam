package backend

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const officialGoRunnerImage = "ghcr.io/01-edu/test-go@sha256:9c57c88e68f8937617a26d805396c775d5a88f43e214b91386986f89a9df13bf"

// RunExercise evaluates code with the same public test image as the Go Final
// Exam. The image owns the test suite; this repository only prepares the
// student workspace that its entrypoint expects at /go-tests/student.
func RunExercise(ex Exercise, code string) (bool, string) {
	studentDir, err := os.MkdirTemp(rootDir, "goexam-*")
	if err != nil {
		return false, "Erreur système: " + err.Error()
	}
	defer os.RemoveAll(studentDir)
	_ = os.Chmod(studentDir, 0o777)

	if err := prepareOfficialGoStudentDir(studentDir, ex.ID, code); err != nil {
		return false, "Erreur de préparation du code: " + err.Error()
	}

	exerciseID := strings.TrimSpace(ex.Runner)
	if exerciseID == "" {
		exerciseID = ex.ID
	}
	stdout, stderr, runErr := RunOfficialGoDocker(studentDir, exerciseID)
	out := strings.TrimSpace(stdout + "\n" + stderr)

	if runErr != nil {
		if out == "" {
			return false, "Échec moulinette Go officielle: " + runErr.Error()
		}
		return false, out
	}
	if out == "" {
		return true, "OK: tous les tests sont validés."
	}
	return true, out
}

func prepareOfficialGoStudentDir(studentDir, exerciseID, code string) error {
	module := `module student

go 1.18

require github.com/01-edu/z01 v0.1.0
`
	if err := os.WriteFile(filepath.Join(studentDir, "go.mod"), []byte(module), 0o644); err != nil {
		return fmt.Errorf("création go.mod: %w", err)
	}

	// The browser editor may submit its last line without a trailing newline.
	// Materialize it as a normal Go source file so the official gofumpt check
	// evaluates formatting, not the transport artifact.
	code = ensureTrailingNewline(code)

	functionSource := RewritePackage(code, "student")
	if err := os.WriteFile(filepath.Join(studentDir, "student.go"), []byte(functionSource), 0o644); err != nil {
		return fmt.Errorf("création student.go: %w", err)
	}

	programDir := filepath.Join(studentDir, filepath.Base(exerciseID))
	if err := os.MkdirAll(programDir, 0o755); err != nil {
		return fmt.Errorf("création dossier exercice: %w", err)
	}
	programSource := RewritePackage(code, "main")
	if err := os.WriteFile(filepath.Join(programDir, "main.go"), []byte(programSource), 0o644); err != nil {
		return fmt.Errorf("création main.go: %w", err)
	}
	return nil
}

func ensureTrailingNewline(code string) string {
	if strings.HasSuffix(code, "\n") {
		return code
	}
	return code + "\n"
}

func RunOfficialGoDocker(studentDir, exerciseID string) (string, string, error) {
	timeout := runTimeout()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	args := officialGoDockerArgs(studentDir, exerciseID)

	cmd := exec.CommandContext(ctx, "docker", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return stdout.String(), stderr.String(), context.DeadlineExceeded
	}
	return stdout.String(), stderr.String(), err
}

func officialGoDockerArgs(studentDir, exerciseID string) []string {
	args := []string{"run", "--rm", "--platform", "linux/amd64"}
	args = append(args, baseDockerRunArgs("")...)
	args = withOptionalCPULimit(args)
	return append(args,
		"-v", studentDir+":/go-tests/student:rw",
		// The official go-tests module resolves `student` as ../piscine-go
		// after copying itself into root's home directory.
		"-v", studentDir+":/root/piscine-go:rw",
		"-e", "EXERCISE="+filepath.Base(exerciseID),
		goRunnerImage(),
	)
}

func goRunnerImage() string {
	return getenvDefault("RUNNER_GO_DOCKER_IMAGE", officialGoRunnerImage)
}
