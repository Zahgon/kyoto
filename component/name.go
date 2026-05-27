package component

// Name implements component name getter/setter,
// required for each component type.
type Name struct {
	Name string
}

// SetName is a component name setter.
func (c *Name) SetName(name string) {
	_ = "STUB: not implemented"

	// GetName is a component name getter.
	return
}

func (c *Name) GetName() string { _ = "STUB: not implemented"; return "" }
