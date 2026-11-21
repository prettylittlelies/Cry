package testkit

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func FixtureDir() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "..", "fixtures")
}

func Fixture(t testing.TB, parts ...string) []byte {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(append([]string{FixtureDir()}, parts...)...))
	if err != nil {
		t.Fatalf("fixture %v: %v", parts, err)
	}
	return b
}
