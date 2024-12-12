package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

/*
To wrap the Logger function into a struct, we can encapsulate the logger 
configuration and the middleware setup into a struct with methods for 
initialization and setup. 
*/

// AppLogger wraps the zap logger and middleware configuration
type AppLogger struct {
	logger *zap.Logger
	config zap.Config
}

// NewAppLogger initializes and returns a new AppLogger instance
func NewAppLogger() (*AppLogger, error) {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		DisableStacktrace: true,
	}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	return &AppLogger{
		logger: logger,
		config: config,
	}, nil
}

// AttachMiddleware attaches the logging middleware to the Echo instance
func (al *AppLogger) AttachMiddleware(e *echo.Echo) {
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
				al.logger.Error("REQUEST",
					zap.String("URI", v.URI),
					zap.Int("STATUS", v.Status),
					logMessage,
					zap.String("METHOD", c.Request().Method),
					zap.String("SESSION_ID", sessionID),
					zap.String("SESSION_TOKEN", sessionToken),
				)
			} else if v.Status >= 200 && v.Status < 300 {
				al.logger.Debug("REQUEST",
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

func (al AppLogger) Error(msg string, fields ...zap.Field) {
	switch len(fields) {
	case 0:
		// No fields provided; log the message alone
		al.logger.Error(msg)
	default:
		// Fields are provided; log the message with fields
		al.logger.Error(msg, fields...)
	}
}

func (al AppLogger) Fatal(msg string, fields ...zap.Field) {
	switch len(fields) {
	case 0:
		al.logger.Fatal(msg)
	default:
		al.logger.Fatal(msg, fields...)
	}
}

func (al AppLogger) Info(msg string, fields ...zap.Field) {
	switch len(fields) {
	case 0:
		al.logger.Info(msg)
	default:
		al.logger.Info(msg, fields...)
	}
}

func (al AppLogger) DPanic(msg string, fields ...zap.Field) {
	switch len(fields) {
	case 0:
		al.logger.DPanic(msg)
	default:
		al.logger.DPanic(msg, fields...)
	}
}

func (al AppLogger) Debug(msg string, fields ...zap.Field) {
	switch len(fields) {
	case 0:
		al.logger.Debug(msg)
	default:
		al.logger.Debug(msg, fields...)
	}
}


