package component

import (
	"time"
)

// Server is a default server component state implementation.
// It uses local temporary files and JSON encoding
// to store, marshal and unmarshal the state.
// Please, make sure this strategy actually fits to your environment.
type Server struct {
	Name

	Path    string        // Path to store component state (default "/tmp/")
	Timeout time.Duration // State timeout (default 24 hours, clean up running on each unmarshal)
}

// path wraps `Path` and resolves with default option.
func (s *Server) path() string { _ = "STUB: not implemented"; return "" }

// timeout wraps `Timeout` and resolves with default option.
func (s *Server) timeout() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// cleanup removes outdated state files.
func (s *Server) cleanup() { _ = "STUB: not implemented"; return }

// Path may be removed while async cleanup is still running.

// Pass if file is not a component state

// File may disappear between ReadDir and Info.

// If creation/modification date is out of timeout bounds,
// remove that file.

// Ignore remove errors to keep cleanup best-effort and panic-free.

// Marshal encodes state with json into temporary file in `Path` directory.
func (s *Server) Marshal(src any) string {
	_ = "STUB: not implemented"
	// Create new tmp file
	return ""
}

// Encode state

// Return filename as a marshaled state

// Unmarshal decodes state with json from temporary file in `Path` directory.
// Fires up a cleanup goroutine in the end.
func (s *Server) Unmarshal(dst any, str string) {
	_ = "STUB: not implemented"
	// Open tmp file
	return
}

// Decode

// Fire up cleanup
