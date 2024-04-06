package v1

import (
	_ "github.com/kldd0/travel-hack-2024/docs"

	"github.com/kldd0/travel-hack-2024/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(handler *echo.Echo, services *service.Services) {
	handler.Use(middleware.Recover())
	handler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: setLogsFile(),
	}))
	handler.Use(DefaultRequestLogger())

	handler.Use(middleware.CORS())

	handler.GET("/health", func(c echo.Context) error { return c.NoContent(200) })
	handler.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := handler.Group("/api/v1")
	{
		newTourRoutes(v1.Group("/tours"), services.Tour, services.Review, services.Order)
		newCityRoutes(v1.Group("/cities"), services.City)
	}
}
