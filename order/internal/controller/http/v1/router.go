package orderV1API

import (
	"github.com/azoma13/computer-assembly-service/order/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(handler *echo.Echo, services *service.Services) {
	handler.Use(middleware.Recover())

	handler.GET("/health", func(c echo.Context) error { return c.NoContent(200) })
	handler.GET("/swagger/*", echoSwagger.WrapHandler)
}
