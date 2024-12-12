package renderer

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/salehrashid/mini-project/internal/app/app_mini_project/util"
	model "github.com/salehrashid/mini-project/pkg/model"

)

// TemplateRenderer is a custom renderer for Echo that uses Go's html/template package.
// It compiles templates and renders them with additional functions registered via funcMap.
type TemplateRenderer struct {
    // Template is the parsed template object containing all available HTML templates.
    Template *template.Template
}

// NewTemplateRenderer initializes a TemplateRenderer by parsing all templates in the directory.
// It registers custom functions for use in templates via funcMap.
// Panics if the templates cannot be parsed.
func NewTemplateRenderer() *TemplateRenderer {
    return &TemplateRenderer{
        Template: template.Must(template.New("t").Funcs(funcMap()).ParseGlob(getTemplateDirectory())),
    }
}

// Render writes the rendered template to the provided writer using the given data.
// Implements the echo.Renderer interface to integrate with the Echo framework.
// Parameters:
// - w: The writer where the rendered HTML will be sent.
// - name: The name of the template to render.
// - data: The data passed to the template for dynamic rendering.
// - c: The Echo context (currently unused in this implementation).
func (renderer *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return renderer.Template.ExecuteTemplate(w, name, data)
}

// funcMap defines custom helper functions to be used in templates (HTML Files).
// These functions can process data before rendering.
// Example: getPageTitle extracts the title from a TemplateModel object.
func funcMap() template.FuncMap {
    return template.FuncMap{
        "getPageTitle": func(templateData model.TemplateModel) string {
            return templateData.Title
        },
    }
}

// getTemplateDirectory retrieves the directory for template files.
// The directory can be set via the environment variable WEB_MINI_PROJECT_TEMPLATE_PATH.
// Defaults to "../../../web/template/html/*/*.html" if the variable is unset.
func getTemplateDirectory() string {
    return util.GetString("WEB_MINI_PROJECT_TEMPLATE_PATH", "../../../web/template/html/*/*.html")
}
