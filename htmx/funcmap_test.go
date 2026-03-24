package htmx

import (
	"html/template"
	"strings"
	"testing"

	"github.com/yznts/kyoto/v3/component"
)

type hxState struct {
	component.Name
}

func (s *hxState) Marshal(src any) string {
	return "encoded-state"
}

func (s *hxState) Unmarshal(dst any, str string) {}

func TestHxStateFunc(t *testing.T) {
	fn, ok := FuncMap["hxstate"].(func(any) template.HTML)
	if !ok {
		t.Fatal("hxstate func has unexpected signature")
	}

	state := &hxState{}
	str := string(fn(state))

	if !strings.Contains(str, `name="hx-state"`) {
		t.Fatalf("unexpected hxstate output: %q", str)
	}
	if !strings.Contains(str, `value="encoded-state"`) {
		t.Fatalf("unexpected hxstate output: %q", str)
	}
}
