package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	model "github.com/salehrashid/mini-project/pkg/model"
)

type LoginHTTPTransport interface {
	RouterRegister(e *echo.Echo)
}

type loginHttpTransport struct{}

func MakeLoginHTTPTransport() LoginHTTPTransport {	
	return loginHttpTransport{}
}

func (t loginHttpTransport) RouterRegister(e *echo.Echo) {
	e.GET("/login", t.login)
}

func (t loginHttpTransport) login(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", echo.Map{
		"TemplateData": model.TemplateModel{
			Title: "Login",
		},
	})
}
