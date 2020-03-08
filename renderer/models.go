package renderer

import "github.com/flosch/pongo2"

type(
	// Renderer : Custom Renderer for templates
	Renderer struct {Debug bool}
)

var (
	ctx pongo2.Context
	t *pongo2.Template
	err error
)
 
