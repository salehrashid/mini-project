package main

import (
	"github.com/labstack/echo/v4"
	"github.com/salehrashid/mini-project/internal/renderer"
	"github.com/salehrashid/mini-project/internal/transport"
	util "github.com/salehrashid/mini-project/internal/util"
)

func main() {
	e := echo.New()
	e.Renderer = renderer.NewTemplateRenderer(true)
	e.Static("/assets", util.GetString("WEB_MINI_PROJECT_FILE_STATICS_PATH", "../../../web/static"))

	transport.Router(e)

	e.Logger.Fatal(e.Start(util.GetPort("DOCKER_WEB_MINI_PROJECT_HOST_PORT")))
}
