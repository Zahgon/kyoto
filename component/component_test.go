package component

import "testing"

func testNamedComponent(_ *Context) State {
	return &Disposable{}
}

func TestComponentGetName(t *testing.T) {
	c := Component(testNamedComponent)
	if got := c.GetName(); got != "testNamedComponent" {
		t.Fatalf("unexpected component name: %q", got)
	}
}
