package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	model "github.com/salehrashid/mini-project/pkg/model"
)

type UserHTTPTransport interface {
	RouterRegister(e *echo.Echo)
}

type userhttpTransport struct{}

func MakeUserHTTPTransport() UserHTTPTransport {
	return userhttpTransport{}
}

func (t userhttpTransport) RouterRegister(e *echo.Echo) {
	e.GET("/home", t.root)
}

func (t userhttpTransport) root(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", echo.Map{
		"Template": model.TemplateModel{
			Title: "Home",
		},
	})
}
