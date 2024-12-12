package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func Logger(e *echo.Echo) {
	// Custom zap configuration
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel), // Set log level
		Development:      true,
		Encoding:         "json", // Log format: JSON or console
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},   // Log to console
		ErrorOutputPaths: []string{"stderr"},   // Error output to console
		DisableStacktrace: true,                // Disable stack traces globally
	}

	// Build the logger
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	// Middleware configuration
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogError:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			sessionID := c.Request().Header.Get("Session-ID")
			sessionToken := c.Request().Header.Get("Session-Token")
			if cookie, err := c.Cookie("session_id"); err == nil {
				sessionID = cookie.Value
			}
			if cookie, err := c.Cookie("session_token"); err == nil {
				sessionToken = cookie.Value
			}

			// Log message dynamically
			logMessage := zap.String("STATUS_MESSAGE", http.StatusText(v.Status))
			if v.Status >= 400 && v.Status < 600 {
				logger.Error("REQUEST",
					zap.String("URI", v.URI),
					zap.Int("STATUS", v.Status),
					logMessage,
					zap.String("METHOD", c.Request().Method),
					zap.String("SESSION_ID", sessionID),
					zap.String("SESSION_TOKEN", sessionToken),
				)
			} else if v.Status >= 200 && v.Status < 300 {
				logger.Debug("REQUEST",
					zap.String("URI", v.URI),
					zap.Int("STATUS", v.Status),
					logMessage,
					zap.String("METHOD", c.Request().Method),
					zap.String("SESSION_ID", sessionID),
					zap.String("SESSION_TOKEN", sessionToken),
				)
			}
			return nil
		},
	}))
}