package component

import "testing"

func TestComponentFuncMapMarshal(t *testing.T) {
	fn, ok := FuncMap["marshal"].(func(State) string)
	if !ok {
		t.Fatal("marshal func has unexpected signature")
	}

	state := &Disposable{}
	if got := fn(state); got != "disposable" {
		t.Fatalf("unexpected marshal output: %q", got)
	}
}
