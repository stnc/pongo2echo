# pongorenderer ECHO V4
Minimal Pongo2 renderer for Echo that just works.

# Usage  
## Installation  

`go get "https://github.com/stnc/pongo-renderer-echo4/renderer"`

The next step is to define renderer and call it

```go
func main() {
  MainRenderer := renderer.Renderer{Debug:true}
  
  // Echo instance
  e := echo.New()
  
  //Use renderer
  e.Renderer = MainRenderer
}

```
somewhere in the program use `Render`  
```go
func index(c echo.Context) error {
	data := pongo2.Context{}
	return c.Render(http.StatusOK, "page.html", data)
}

```

thanks to https://github.com/siredwin/pongorenderer (for echo framework v3)
