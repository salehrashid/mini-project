package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/salehrashid/mini-project/internal/renderer"
	"github.com/salehrashid/mini-project/internal/transport"
	util "github.com/salehrashid/mini-project/internal/util"
	logger "github.com/salehrashid/mini-project/pkg/util"
)

func main() {
	// Initialize a new Echo instance for HTTP server setup.
	e := echo.New()

	// Set the template renderer to render HTML pages.
	// The renderer is responsible for locating and parsing the HTML templates.
	e.Renderer = renderer.NewTemplateRenderer()

	// Serve static files from the specified path.
	// The "/assets" route maps to the static files directory.
	e.Static("/assets", util.GetString("WEB_MINI_PROJECT_FILE_STATICS_PATH", "../../../web/static"))

	// Add middleware to log all HTTP requests for debugging and analysis.
	logger.Logger(e)

	// Add middleware to recover from panics and log the errors.
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	// Initialize the HTTP transport layer and register the routes.
	// The transport handles route definitions and their corresponding handlers.
	transport := transport.MakeHTTPTransport()
	transport.RouterRegister(e)

	// Start the HTTP server on the specified port and log any fatal errors.
	e.Logger.Fatal(e.Start(util.GetPort("DOCKER_WEB_MINI_PROJECT_HOST_PORT")))
}
