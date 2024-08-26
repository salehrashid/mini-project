package renderer

import (
	"github.com/labstack/echo/v4"
	"github.com/salehrashid/mini-project/internal/util"
	"html/template"
	"io"
)

type Renderer struct {
	Template *template.Template
	Debug    bool
	Location string
}

type DynamicTitle struct {
	Title string
}

func NewTemplateRenderer(debug bool) *Renderer {
	t := new(Renderer)
	t.Debug = debug

	t.ReloadTemplates()

	return t
}

func (renderer *Renderer) ReloadTemplates() {
	renderer.Template = template.Must(template.ParseGlob(getTemplateDirectory()))
}

func (renderer *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if renderer.Debug {
		renderer.ReloadTemplates()
	}
	return renderer.Template.ExecuteTemplate(w, name, data)
}

func getTemplateDirectory() string {
	return util.GetString("WEB_MINI_PROJECT_TEMPLATE_PATH", "../../../web/template/html/*/*.html")
}
