package rendering

import (
	"html/template"
	"io"
	"testing"

	"github.com/yznts/kyoto/v3/component"
)

type inlineRendererState struct {
	component.Name
	body string
}

func (*inlineRendererState) Marshal(src any) string        { return "" }
func (*inlineRendererState) Unmarshal(dst any, str string) {}
func (*inlineRendererState) RenderSkip() bool              { return false }
func (s *inlineRendererState) Render(_ component.State, out io.Writer) error {
	_, err := io.WriteString(out, s.body)
	return err
}

type inlinePlainState struct {
	component.Universal
	Value string `json:"value"`
}

func TestRenderFuncMapRendersFuture(t *testing.T) {
	fn, ok := FuncMap["render"].(func(component.Future) template.HTML)
	if !ok {
		t.Fatal("render func has unexpected signature")
	}

	state := &inlineRendererState{body: "inline"}
	out := string(fn(func() component.State { return state }))
	if out != "inline" {
		t.Fatalf("unexpected inline render output: %q", out)
	}
}

func TestRenderFuncMapPanicsForNonRenderer(t *testing.T) {
	fn := FuncMap["render"].(func(component.Future) template.HTML)

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for non-renderer in render func")
		}
	}()

	_ = fn(func() component.State { return &inlinePlainState{Value: "x"} })
}

func TestFuncMapAllContainsAllFunctions(t *testing.T) {
	if _, ok := FuncMapAll["render"]; !ok {
		t.Fatal("render function is missing from merged funcmap")
	}
	if _, ok := FuncMapAll["hxstate"]; !ok {
		t.Fatal("hxstate function is missing from merged funcmap")
	}
	if _, ok := FuncMapAll["marshal"]; !ok {
		t.Fatal("marshal function is missing from merged funcmap")
	}
}
