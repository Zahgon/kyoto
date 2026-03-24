package component

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestServerMarshalUnmarshal(t *testing.T) {
	dir := t.TempDir()
	server := &Server{Path: dir}

	state := map[string]string{"k": "v"}
	file := server.Marshal(state)

	if !strings.HasSuffix(file, ".component") {
		t.Fatalf("unexpected state file name: %q", file)
	}
	if _, err := os.Stat(filepath.Join(dir, file)); err != nil {
		t.Fatalf("state file was not created: %v", err)
	}

	var decoded map[string]string
	server.Unmarshal(&decoded, file)
	if got, want := decoded["k"], "v"; got != want {
		t.Fatalf("server unmarshal mismatch: got %q, want %q", got, want)
	}
}

func TestServerCleanup(t *testing.T) {
	dir := t.TempDir()
	// Timeout of 1 hour so the freshly-created file is kept, but the
	// backdated file is expired.
	server := &Server{Path: dir, Timeout: time.Hour}

	freshFile := filepath.Join(dir, "fresh.component")
	if err := os.WriteFile(freshFile, []byte("{}"), 0o600); err != nil {
		t.Fatalf("failed to create fresh state file: %v", err)
	}

	oldFile := filepath.Join(dir, "old.component")
	if err := os.WriteFile(oldFile, []byte("{}"), 0o600); err != nil {
		t.Fatalf("failed to create old state file: %v", err)
	}
	oldTime := time.Now().Add(-2 * time.Hour)
	if err := os.Chtimes(oldFile, oldTime, oldTime); err != nil {
		t.Fatalf("failed to set old modtime: %v", err)
	}

	server.cleanup()

	if _, err := os.Stat(oldFile); !os.IsNotExist(err) {
		t.Fatalf("expected expired state file to be removed, stat err=%v", err)
	}
	if _, err := os.Stat(freshFile); err != nil {
		t.Fatalf("expected fresh state file to be kept: %v", err)
	}
}
