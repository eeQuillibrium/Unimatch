package templates

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	template *template.Template
}

func NewTemplate(assetsPath string) *Template {
	return &Template{template: template.Must(template.ParseGlob(assetsPath + "/*.html"))}
}
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.template.ExecuteTemplate(w, name, data)
}
