package component

import "testing"

func TestDisposableMarshalUnmarshal(t *testing.T) {
	d := &Disposable{}
	if got := d.Marshal(struct{}{}); got != "disposable" {
		t.Fatalf("unexpected marshal result: %q", got)
	}

	before := 42
	after := before
	d.Unmarshal(&after, "ignored")
	if after != before {
		t.Fatal("disposable unmarshal must be no-op")
	}
}
