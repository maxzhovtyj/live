package template

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func New() echo.Renderer {
	return &Template{
		Templates: template.Must(template.ParseGlob("internal/public/*.html")),
	}
}
