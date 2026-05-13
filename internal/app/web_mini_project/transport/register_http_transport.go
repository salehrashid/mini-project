package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	model "github.com/salehrashid/mini-project/pkg/model"
)

type RegisterHTTPTransport interface {
	RouterRegister(e *echo.Echo)
}

type registerHttpTransport struct{}

func MakeRegisterHTTPTransport() RegisterHTTPTransport {
	return registerHttpTransport{}
}

func (t registerHttpTransport) RouterRegister(e *echo.Echo) {
	e.GET("/register", t.register)
}

func (t registerHttpTransport) register(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", echo.Map{
		"TemplateData": model.TemplateModel{
			Title: "Register",
		},
	})
}
