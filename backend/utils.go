package backend

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func DirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func FileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func Trim(s string) string {
	return strings.TrimSpace(s)
}

// RewritePackage keeps the editor flexible while producing the package layout
// required by the official 01 Edu image.
func RewritePackage(src, newPkg string) string {
	lines := strings.Split(src, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if strings.HasPrefix(trimmed, "package ") {
			lines[i] = "package " + newPkg
			return strings.Join(lines, "\n")
		}
		break
	}
	return "package " + newPkg + "\n\n" + src
}

func runTimeout() time.Duration {
	raw := getenvDefault("RUNNER_TIMEOUT_SECONDS", "10")
	sec, err := strconv.Atoi(raw)
	if err != nil || sec <= 0 {
		return 10 * time.Second
	}
	return time.Duration(sec) * time.Second
}

func getenvDefault(key, def string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return def
}
