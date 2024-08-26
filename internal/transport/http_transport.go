package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salehrashid/mini-project/internal/renderer"
)

type HTTPTransport interface {
	Router(e *echo.Echo)
}

type httpTransport struct{}

func NewHTTPTransport() HTTPTransport {
	return httpTransport{}
}

func (t httpTransport) Router(e *echo.Echo) {
	e.GET("/", t.root)
}

func (t httpTransport) root(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", echo.Map{
		"Template": renderer.TemplateModel{
			Title: "Home",
		},
	})
}
