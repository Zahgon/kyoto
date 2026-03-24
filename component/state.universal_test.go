package component

import (
	"strings"
	"testing"
)

type universalState struct {
	Universal
	Value string `json:"value"`
}

func TestUniversalRoundTrip(t *testing.T) {
	state := &universalState{Value: "spaces/and?symbols=ok"}

	encoded := state.Marshal(state)
	if strings.Contains(encoded, " ") {
		t.Fatal("marshal output should be safe for attribute usage")
	}

	decoded := &universalState{}
	state.Unmarshal(decoded, encoded)
	if got, want := decoded.Value, state.Value; got != want {
		t.Fatalf("roundtrip mismatch: got %q, want %q", got, want)
	}
}
