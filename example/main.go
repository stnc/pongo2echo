package main

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
	"github.com/stnc/pongo2echo"
)

var (
	data         = pongo2.Context{}
	mainRenderer = pongo2echo.Renderer{Debug: true} // use any renderer
)

//GetAllData all list
func GetAllData(c echo.Context) error {

	posts := []string{
		"Larry Ellison",
		"Carlos Slim Helu",
		"Mark Zuckerberg",
		"Amancio Ortega ",
		"Jeff Bezos",
		" Warren Buffet ",
		"Bill Gates",
		"selman tunç",
	}

	return c.Render(http.StatusOK, "templates/index.html",
		pongo2.Context{"title": "hello echo fw", "posts": posts})

}

func main() {
	e := echo.New()
	e.Renderer = mainRenderer //pongo2 init
	// http://localhost:8888/home
	e.GET("/home", GetAllData)
	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}
