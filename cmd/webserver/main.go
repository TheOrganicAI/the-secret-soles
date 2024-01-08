package main

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type TemplateRegistry struct {
	templates *template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	templatesPath := "./templates/*.html"
	templateRegistry := &TemplateRegistry{
		templates: template.Must(template.ParseGlob(templatesPath)),
	}
	e.Renderer = templateRegistry

	e.Static("/assets", "./assets")

	e.GET("/", homeHandler)
	e.GET("/contact", contactHandler)
	e.GET("/gallery", galleryHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

func homeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func contactHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "contact.html", nil)
}

func galleryHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "gallery.html", nil)
}
