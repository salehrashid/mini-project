package renderer

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/salehrashid/mini-project/internal/util"
	model "github.com/salehrashid/mini-project/pkg/model"

)

type TemplateRenderer struct {
	Template *template.Template
	Location string
}

func NewTemplateRenderer() *TemplateRenderer {
	return &TemplateRenderer{
		Template: template.Must(template.New("t").Funcs(funcMap()).ParseGlob(getTemplateDirectory())),
	}
}

func (renderer *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return renderer.Template.ExecuteTemplate(w, name, data)
}

//variable passer to template / html
func funcMap() template.FuncMap{
	return template.FuncMap{
		"getPageTitle": func(templateData model.TemplateModel) string {
			return templateData.Title
		},
	}
}

func getTemplateDirectory() string {
	return util.GetString("WEB_MINI_PROJECT_TEMPLATE_PATH", "../../../web/template/html/*/*.html")
}
