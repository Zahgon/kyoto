package component

import (
	"net/http/httptest"
	"testing"
)

func TestNewContextAndMapStore(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	ctx := NewContext(rec, req)

	if ctx.ResponseWriter != rec {
		t.Fatal("response writer is not wired")
	}
	if ctx.Request != req {
		t.Fatal("request is not wired")
	}

	ctx.Set("k", "v")
	if got := ctx.Get("k"); got != "v" {
		t.Fatalf("unexpected store value: %v", got)
	}
}
