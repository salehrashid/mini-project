package util

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func Logger(e *echo.Echo) {
	logger, _ := zap.NewDevelopment()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogError:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			switch v.Status {
			case 404:
				logger.Error("REQUEST",
					zap.String("URI", v.URI),
					zap.Int("STATUS", v.Status),
				)
			case 200:
				logger.Debug("REQUEST",
					zap.String("URI", v.URI),
					zap.Int("STATUS", v.Status),
				)
			}
			return nil
		},
	}))
}
