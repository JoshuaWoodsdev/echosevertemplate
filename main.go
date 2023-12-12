package main

import (
    "html/template"
    "net/http"
    "io"
    "github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom renderer for rendering HTML templates
type TemplateRenderer struct {
    templates *template.Template
}

// Render implements the echo.Renderer interface
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
    // Initialize Echo
    e := echo.New()

    // Load HTML templates
    renderer := &TemplateRenderer{
        templates: template.Must(template.ParseGlob("templates/*.html")),
    }
    e.Renderer = renderer

    // Define routes
    e.GET("/", func(c echo.Context) error {
        return c.Render(http.StatusOK, "index.html", map[string]interface{}{"title": "Index Page"})
    })

    e.GET("/about", func(c echo.Context) error {
        return c.Render(http.StatusOK, "about.html", map[string]interface{}{"title": "About Page"})
    })

    // Start the server
    e.Start(":8080")
}
