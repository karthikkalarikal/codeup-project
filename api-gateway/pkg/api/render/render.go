package render

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	API             string
	CSSVersion      string
}

type TemplateRenderer struct {
	Templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	fmt.Println("name", name, "data", data)
	return t.Templates.ExecuteTemplate(w, name, data)
}
