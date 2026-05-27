package rendering

import (
	"net/http"

	"github.com/yznts/kyoto/v3/component"
)

// Handler builds a http.HandlerFunc that renders provided component.
func Handler(c component.Component) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// Create context

// Build page state tree

// Inject component name, unless it's already set

// Ensure state implements render

// Check if we need to skip rendering

// Render
