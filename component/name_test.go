package component

import "testing"

func TestNameGetSet(t *testing.T) {
	n := &Name{}
	n.SetName("foo")
	if got := n.GetName(); got != "foo" {
		t.Fatalf("unexpected name: %q", got)
	}
}
