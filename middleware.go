package pongo2echo

import (
	"errors"
	"io"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
)

type (
	// Renderer : Custom Renderer for templates
	Renderer struct{ Debug bool }
)

var (
	ctx pongo2.Context
	t   *pongo2.Template
	err error
)

// Render : Custom renderer
func (r Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if data != nil {
		var ok bool
		ctx, ok = data.(pongo2.Context)

		if !ok {
			return errors.New("no pongo2.Context data was passed")
		}
	}

	if r.Debug {
		t, err = pongo2.FromFile(name)
	} else {
		t, err = pongo2.FromCache(name)
	}
	// Add some static values
	ctx["version_number"] = "v3.0"
	if err != nil {
		return err
	}
	return t.ExecuteWriter(ctx, w)
}
