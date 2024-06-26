package v1

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func DefaultRequestLogger() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogHost:    true,
		LogMethod:  true,
		LogError:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			statusCode := v.Status
			var e *zerolog.Event
			if v.Error != nil {
				e = log.Error().Err(v.Error)
			} else {
				e = log.Info()
			}

			e.
				Str("method", v.Method).
				Str("host", v.Host).
				Str("URI", v.URI).
				Int("status", statusCode).
				Str("latency", v.Latency.String()).
				Msg("incoming request")

			return nil
		},
	})
}

func setLogsFile() *os.File {
	file, err := os.OpenFile("logs/requests.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	return file
}
