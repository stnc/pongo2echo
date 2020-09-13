Pongo2gin
=========

Package pongo2echo is a template renderer that can be used with the minimalist Go web framework
https://github.com/labstack/echo it uses the Pongo2 template library
https://github.com/flosch/pongo2

## Installation  

`go get "github.com/stnc/pongo2echo"`

Requirements
------------

Requires Echo 4+ or higher and Pongo2.

Usage
-----

To use pongo2echo you need to set your router.HTMLRenderer to a new renderer
instance, this is done after creating the Gin router when the Gin application
starts up. This assumes templates will be located in the "templates"
directory, or you can use pongo2gin.TemplatePath("templates") to specify a custom location.

To render templates from a route, call c.HTML just as you would with
regular Gin templates, the only difference is that you pass template
data as a pongo2.Context instead of gin.H type.


![Screen](https://raw.githubusercontent.com/stnc/pongo2gin/master/example/ginScreen.png)

Basic Example
-------------

```go
package main
import (
	"log"
	"net/http"
	"github.com/stnc/pongo2gin"
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)
//GetAllData all list
func GetAllData(c *gin.Context) {
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
	// Call the HTML method of the Context to render a template
	c.HTML(http.StatusOK, "index.html",
		pongo2.Context{
			"title": "hello pongo",
			"posts": posts,
		},
	)
}

func main() {

	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(gin.Recovery())
	r.HTMLRender = pongo2gin.TemplatePath("templates")
	r.GET("home", GetAllData)
	log.Fatal(r.Run(":8888"))
}
```

HTML 
----------------


```html

<h1> {{ title }}</h1>

{% for post in posts%}
<ul>
  <li>{{post}}</li>
</ul>
{% endfor %}

```


GoDoc
-----


Specials Thanks
-----

https://github.com/siredwin/pongorenderer Version 3
https://machiel.me/post/pongo2-with-echo-or-net-http/