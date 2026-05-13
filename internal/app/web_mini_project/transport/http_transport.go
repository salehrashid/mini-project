package transport

import (
	"net/http"

	"github.com/labstack/echo/v4"
	model "github.com/salehrashid/mini-project/pkg/model"
	internalTransport "github.com/salehrashid/mini-project/internal/pkg/transport"
	publicUtil "github.com/salehrashid/mini-project/pkg/util"
)

// HTTPTransport is an interface defining the contract for HTTP transport layers.
// It requires implementing a method to register routes with the Echo framework.
type HTTPTransport interface {
	// RouterRegister is used to define and register HTTP routes.
	// The method takes an instance of Echo as a parameter for route setup.
	RouterRegister(e *echo.Echo)

	GetErrorHTTPTransport() internalTransport.ErrorHTTPTransport
}

// httpTransport is a concrete implementation of the HTTPTransport interface.
// It embeds UserHTTPTransport and etc, allowing delegation of some functionality.
type httpTransport struct {
	publicUtil.AppLogger

	UserHTTPTransport // Embedded struct to handle user-specific routes.
	LoginHTTPTransport
	RegisterHTTPTransport
	internalTransport.ErrorHTTPTransport
}

// MakeHTTPTransport initializes and returns an instance of HTTPTransport.
// It creates an httpTransport with the required dependencies.
func MakeHTTPTransport(logger publicUtil.AppLogger) HTTPTransport {
	return httpTransport{
		ErrorHTTPTransport: internalTransport.MakeErrorHTTPTransport(logger),

		UserHTTPTransport: MakeUserHTTPTransport(),
		LoginHTTPTransport: MakeLoginHTTPTransport(),
		RegisterHTTPTransport: MakeRegisterHTTPTransport(),
	}
}

// RouterRegister registers HTTP routes for the httpTransport.
// It sets up the dashboard route and delegates additional route registration to UserHTTPTransport and etc.
func (t httpTransport) RouterRegister(e *echo.Echo) {
	// Register the dashboard route that serves the home page.
	e.GET("/dashboard", t.dashboard)
	e.GET("/", t.root)

	// Delegate route registration to the embedded UserHTTPTransport and etc.
	t.UserHTTPTransport.RouterRegister(e)
	t.LoginHTTPTransport.RouterRegister(e)
	t.RegisterHTTPTransport.RouterRegister(e)
}

// dashboard is the handler function for the dashboard route ("/").
// It renders the "dashboard.html" template with a title passed as data.
func (t httpTransport) dashboard(c echo.Context) error {
	return c.Render(http.StatusOK, "dashboard.html", echo.Map{
		"TemplateData": model.TemplateModel{
			Title: "Dashboard", // Page title to be displayed in the template.
		},
	})
}

// root is the handler function for the root route ("/").
// It redirects to the dashboard page.
func (t httpTransport) root(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/dashboard")
}

func (t httpTransport) GetErrorHTTPTransport() internalTransport.ErrorHTTPTransport{
	return t.ErrorHTTPTransport
}