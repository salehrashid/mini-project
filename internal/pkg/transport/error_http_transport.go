package transport

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salehrashid/mini-project/pkg/model"
	publicUtil "github.com/salehrashid/mini-project/pkg/util"
	"go.uber.org/zap"
)

// ErrorHTTPTransport defines an interface for handling error-related routing and error page rendering.
type ErrorHTTPTransport interface {
	// RouterRegister registers the error handler with the Echo router.
	RouterRegister(e *echo.Echo)

	// ErrorPageTransport handles rendering of custom error pages.
	// It takes an error and the Echo context as parameters.
	ErrorPageTransport(err error, c echo.Context)
}

// errorHttpTransport implements the ErrorHTTPTransport interface and manages error handling logic.
type errorHttpTransport struct {
	// logger is used for logging errors and debugging information.
	logger publicUtil.AppLogger
}

// MakeErrorHTTPTransport creates a new instance of errorHttpTransport with the provided logger.
// This function returns an instance that adheres to the ErrorHTTPTransport interface.
func MakeErrorHTTPTransport(logger publicUtil.AppLogger) ErrorHTTPTransport {
	return errorHttpTransport{
		logger: logger,
	}
}

// RouterRegister sets up the error handler by assigning ErrorPageTransport
// to the HTTPErrorHandler field of the Echo framework.
func (t errorHttpTransport) RouterRegister(e *echo.Echo) {
	e.HTTPErrorHandler = t.ErrorPageTransport
}

// ErrorPageTransport is a custom error handler that renders an error page based on the error type.
// - If the error is an Echo HTTP error, it extracts the status code and message.
// - For generic errors, it uses the default HTTP 500 status code and the error message.
// - The error page is rendered using the `error.html` template.
// - Logs errors using the logger.
func (t errorHttpTransport) ErrorPageTransport(err error, c echo.Context) {
	var message interface{}
	code := http.StatusInternalServerError

	// Check if the error is an Echo HTTPError and extract details.
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message
	} else {
		message = err.Error()
	}

	// Attempt to render the custom error page with error details.
	err = c.Render(code, "error.html", echo.Map{
		"Code":    code,                  // HTTP status code of the error.
		"Message": message,               // Error message to display.
		"TemplateData": model.TemplateModel{
			Title: fmt.Sprintf("Error %d", code), // Set a dynamic title for the error page.
		},
	})

	// Log rendering errors or the original error.
	if err != nil {
		t.logger.Error("errorHttpTransport.ErrorPageTransport", zap.Error(err))
		return
	}

	// Log the initial error for debugging purposes.
	t.logger.Error("errorHttpTransport.ErrorPageTransport", zap.Error(err))
}
