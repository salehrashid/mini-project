package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salehrashid/mini-project/internal/renderer"
)

func httpTransport(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", renderer.DynamicTitle{Title: "Home"})
}

func Router(e *echo.Echo) {
	e.GET("/", httpTransport)
}
