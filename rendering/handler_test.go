package rendering

import (
	"errors"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/yznts/kyoto/v3/component"
)

type handlerState struct {
	component.Name
	body       string
	skip       bool
	renderErr  error
	renderHits int
}

func (*handlerState) Marshal(src any) string        { return "" }
func (*handlerState) Unmarshal(dst any, str string) {}
func (s *handlerState) RenderSkip() bool            { return s.skip }
func (s *handlerState) Render(_ component.State, out io.Writer) error {
	s.renderHits++
	if s.renderErr != nil {
		return s.renderErr
	}
	_, err := io.WriteString(out, s.body)
	return err
}

type handlerPlainState struct {
	component.Universal
	Value string `json:"value"`
}

func testHandlerComponentFactory(s component.State) component.Component {
	return func(_ *component.Context) component.State {
		return s
	}
}

func TestHandlerRendersAndSetsNameWhenEmpty(t *testing.T) {
	state := &handlerState{body: "ok"}
	h := Handler(testHandlerComponentFactory(state))

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	h(rec, req)

	if got := rec.Body.String(); got != "ok" {
		t.Fatalf("unexpected response body: %q", got)
	}
	if state.GetName() == "" {
		t.Fatal("handler should inject component name when state name is empty")
	}
	if state.renderHits != 1 {
		t.Fatalf("expected single render call, got %d", state.renderHits)
	}
}

func TestHandlerKeepsExistingName(t *testing.T) {
	state := &handlerState{body: "ok"}
	state.SetName("already-set")
	h := Handler(testHandlerComponentFactory(state))

	h(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))

	if got := state.GetName(); got != "already-set" {
		t.Fatalf("handler should keep existing name, got %q", got)
	}
}

func TestHandlerSkipsRender(t *testing.T) {
	state := &handlerState{skip: true, body: "must-not-render"}
	h := Handler(testHandlerComponentFactory(state))

	rec := httptest.NewRecorder()
	h(rec, httptest.NewRequest("GET", "/", nil))

	if rec.Body.Len() != 0 {
		t.Fatalf("expected empty body, got %q", rec.Body.String())
	}
	if state.renderHits != 0 {
		t.Fatalf("expected no render calls, got %d", state.renderHits)
	}
}

func TestHandlerPanicsWhenNotRenderer(t *testing.T) {
	h := Handler(testHandlerComponentFactory(&handlerPlainState{Value: "x"}))

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for non-renderer state")
		}
	}()

	h(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))
}

func TestHandlerPanicsWhenRenderFails(t *testing.T) {
	state := &handlerState{renderErr: errors.New("boom")}
	h := Handler(testHandlerComponentFactory(state))

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic when renderer returns an error")
		}
	}()

	h(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))
}
