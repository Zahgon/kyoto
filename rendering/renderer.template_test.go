package rendering

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/yznts/kyoto/v3/component"
)

type templateState struct {
	component.Universal
	Template
	Value string `json:"value"`
}

func TestTemplateRenderWithRawTemplate(t *testing.T) {
	state := &templateState{Value: "world"}
	state.SetName("raw")
	state.Raw = template.Must(template.New("raw").Parse("hello {{.Value}}"))

	var out bytes.Buffer
	if err := state.Render(state, &out); err != nil {
		t.Fatalf("unexpected render error: %v", err)
	}
	if got := out.String(); got != "hello world" {
		t.Fatalf("unexpected template output: %q", got)
	}
}

func TestTemplateRenderFromDiskGlob(t *testing.T) {
	dir := t.TempDir()
	name := "view.html"
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte("{{.Value}}"), 0o600); err != nil {
		t.Fatalf("failed to write template file: %v", err)
	}

	state := &templateState{Value: "from-disk"}
	state.SetName(name)
	state.Glob = filepath.Join(dir, "*.html")

	var out bytes.Buffer
	if err := state.Render(state, &out); err != nil {
		t.Fatalf("unexpected render error: %v", err)
	}
	if got := strings.TrimSpace(out.String()); got != "from-disk" {
		t.Fatalf("unexpected template output: %q", got)
	}
	if state.FuncMap == nil {
		t.Fatal("template funcmap default should be set")
	}
}
