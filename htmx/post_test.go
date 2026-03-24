package htmx

import (
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/yznts/kyoto/v3/component"
)

type postState struct {
	component.Name
	lastUnmarshal string
	unmarshalHits int
}

func (s *postState) Marshal(src any) string {
	return "encoded-state"
}

func (s *postState) Unmarshal(dst any, str string) {
	s.lastUnmarshal = str
	s.unmarshalHits++
}

func TestPostSkipsOnNonPost(t *testing.T) {
	ctx := component.NewContext(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))
	state := &postState{}
	hit := false

	Post(ctx, state, func() { hit = true })

	if hit {
		t.Fatal("handler should not run on non-POST requests")
	}
	if state.unmarshalHits != 0 {
		t.Fatal("state should not unmarshal on non-POST requests")
	}
}

func TestPostSkipsWhenHxStateMissing(t *testing.T) {
	body := strings.NewReader(url.Values{"foo": {"bar"}}.Encode())
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	ctx := component.NewContext(httptest.NewRecorder(), req)

	state := &postState{}
	hit := false
	Post(ctx, state, func() { hit = true })

	if hit {
		t.Fatal("handler should not run when hx-state is missing")
	}
	if state.unmarshalHits != 0 {
		t.Fatal("state should not unmarshal when hx-state is missing")
	}
}

func TestPostPanicsForDisposableState(t *testing.T) {
	body := strings.NewReader(url.Values{"hx-state": {"disposable"}}.Encode())
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	ctx := component.NewContext(httptest.NewRecorder(), req)

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for disposable state")
		}
	}()

	Post(ctx, &postState{}, func() {})
}

func TestPostUnmarshalsAndCallsHandler(t *testing.T) {
	body := strings.NewReader(url.Values{"hx-state": {"abc123"}}.Encode())
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	ctx := component.NewContext(httptest.NewRecorder(), req)

	state := &postState{}
	hit := false
	Post(ctx, state, func() { hit = true })

	if !hit {
		t.Fatal("handler must run for valid hx-state")
	}
	if state.unmarshalHits != 1 {
		t.Fatalf("expected 1 unmarshal call, got %d", state.unmarshalHits)
	}
	if state.lastUnmarshal != "abc123" {
		t.Fatalf("unexpected unmarshal input: %q", state.lastUnmarshal)
	}
}
