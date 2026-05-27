package rendering

import (
	"embed"
	"html/template"
	"io"

	"github.com/yznts/kyoto/v3/component"
)

// Global template configuration defaults.
// We're providing them to make it easier to configure rendering defaults across the project.
var (
	TEMPLATE_GLOB              = "*.html"
	TEMPLATE_FUNCMAP           = FuncMapAll
	TEMPLATE_EMBEDFS *embed.FS = nil
)

// Template is a html/template renderer.
// Use Raw to provide handmade template,
// or provide template building parameters (Name, Glob, etc.).
type Template struct {
	Raw  *template.Template `json:"-"` // Raw template will be used instead if provided
	Name string             // Resolved from component name by default
	Skip bool               `json:"-"` // false by default

	Glob    string           `json:"-"` // *.html by default
	EmbedFS *embed.FS        `json:"-"` // nil by default
	FuncMap template.FuncMap `json:"-"` // render.FuncMap by default
}

func (t *Template) RenderSkip() bool { _ = "STUB: not implemented"; return false }

func (t *Template) Render(state component.State, w io.Writer) error {
	_ = "STUB: not implemented"
	// Defaults
	return nil
}

// Define template

// Base

// Functions

// Parse

// Parse embedded

// Parse embedded

// Parse from disk

// Render
