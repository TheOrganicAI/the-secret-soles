package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// TemplateRegistry is a custom type that holds a reference to a parsed template
type TemplateRegistry struct {
	templates *template.Template
}

// Render makes TemplateRegistry implement the echo.Renderer interface so you
// can use it as a custom template renderer.
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Create a new instance of Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize the template registry with parsed templates located in the /templates folder
	templatesPath := "./templates/*.html" // Adjust the path according to your main.go file's location
	templateRegistry := &TemplateRegistry{
		templates: template.Must(template.ParseGlob(templatesPath)),
	}
	e.Renderer = templateRegistry

	// Static file serving
	e.Static("/assets", "./assets")

	// Routes
	e.GET("/", homeHandler)
	e.GET("/contact", contactHandler)
	e.GET("/gallery", galleryHandler)

	// Start the server on port 8080
	e.Logger.Fatal(e.Start(":8081"))
}

// Home page handler
func homeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

// Contact page handler
func contactHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "contact.html", nil)
}

// Gallery page handler
func galleryHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "gallery.html", nil)
}
