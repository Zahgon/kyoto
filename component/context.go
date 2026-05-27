package component

import (
	"net/http"
	"sync"
)

// Context is the context of the current request.
// It is passed to the pages and components.
type Context struct {
	// Handler
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	// Store
	Store
}

// Initialize a new context, that will be passed through the components.
// Uses MapStore as a store by default.
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	_ = "STUB: not implemented"
	return nil
}

// Store allows you to store own data inside of the context.
type Store interface {
	Get(key string) any
	Set(key string, value any)
}

type MapStore struct {
	mu    sync.RWMutex
	store map[string]any
}

func (s *MapStore) Get(key string) any { _ = "STUB: not implemented"; return *new(any) }

func (s *MapStore) Set(key string, value any) { _ = "STUB: not implemented"; return }

func NewMapStore() *MapStore { _ = "STUB: not implemented"; return nil }
