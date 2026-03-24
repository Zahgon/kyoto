package component

import (
	"net/http/httptest"
	"testing"
)

func testUseComponent(_ *Context) State {
	return &Disposable{}
}

func TestUseSetsComponentName(t *testing.T) {
	ctx := NewContext(httptest.NewRecorder(), httptest.NewRequest("GET", "/", nil))
	fut := Use(ctx, Component(testUseComponent))
	state := fut()

	if got, want := state.GetName(), "testUseComponent"; got != want {
		t.Fatalf("state name mismatch: got %q, want %q", got, want)
	}
}
