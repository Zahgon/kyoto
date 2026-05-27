package component

// Component represents a component state builder, defined as a function.
type Component func(ctx *Context) State

// GetName returns the name of the component,
// based on the function name.
func (c Component) GetName() string { _ = "STUB: not implemented"; return "" }
