package component

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestServerMarshalUnmarshalAndCleanup(t *testing.T) {
	dir := t.TempDir()
	server := &Server{Path: dir, Timeout: 10 * time.Millisecond}

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

	oldFile := filepath.Join(dir, "old.component")
	if err := os.WriteFile(oldFile, []byte("{}"), 0o600); err != nil {
		t.Fatalf("failed to create old state file: %v", err)
	}
	oldTime := time.Now().Add(-time.Hour)
	if err := os.Chtimes(oldFile, oldTime, oldTime); err != nil {
		t.Fatalf("failed to set old modtime: %v", err)
	}

	server.cleanup()
	if _, err := os.Stat(oldFile); !os.IsNotExist(err) {
		t.Fatalf("expected old state file to be removed, stat err=%v", err)
	}
}
