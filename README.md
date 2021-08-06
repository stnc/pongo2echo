Pongo2echo
=========

Package pongo2echo is a template renderer that can be used with the minimalist Go web framework
https://github.com/labstack/echo it uses the Pongo2 template library
https://github.com/flosch/pongo2

pongo2 is a Django-syntax like templating-language (official website).

## Installation  

`go get "github.com/stnc/pongo2echo"`

## Here is Compatible with pongo version 4
 [Pongo4echo](https://github.com/stnc/pongo4echo) - pongo2 echo framework stability renderer / Compatible with pongo version 4

# please don't forget to give stars :)



Requirements
------------

Requires Echo 4+ or higher and Pongo2.

Usage
-----

Real Example [echo+pongo+gorm+pagination]

https://github.com/stnc/golang-echo-realworld-example-web-app

![Screen](https://raw.githubusercontent.com/stnc/pongo2echo/master/example/echoScreen.png)

Basic Example
-------------

```go
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
	e.Debug = true
	// http://localhost:8888/home
	e.GET("/home", GetAllData)
	// Start server
	e.Logger.Fatal(e.Start(":8888"))
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
