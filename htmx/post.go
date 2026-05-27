package htmx

import "github.com/yznts/kyoto/v3/component"

// Post is a helper function that simplifies the handling of stateful htmx POST requests.
func Post(ctx *component.Context, state component.State, handler func()) {
	_ = "STUB: not implemented"
	// We are only interested in POST requests here
	return
}

// Parse the form to get the state

// If no state is present in the form, we ignore.
// Porbably this is a regular POST request, not related to htmx.

// If the state is disposable, we panic.
// This is a safety measure to prevent misuse of disposable components.

// Unmarshal the state from the form

// Call the handler
