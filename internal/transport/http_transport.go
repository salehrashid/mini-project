package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	model "github.com/salehrashid/mini-project/pkg/model"
)

type HTTPTransport interface {
	RouterRegister(e *echo.Echo)
}

type httpTransport struct {
	UserHTTPTransport
}

func MakeHTTPTransport() HTTPTransport {
	return httpTransport{
		UserHTTPTransport: MakeUserHTTPTransport(),
	}
}

func (t httpTransport) RouterRegister(e *echo.Echo) {
	e.GET("/", t.root)

	t.UserHTTPTransport.RouterRegister(e)
}

func (t httpTransport) root(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", echo.Map{
		"Template": model.TemplateModel{
			Title: "Home",
		},
	})
}
