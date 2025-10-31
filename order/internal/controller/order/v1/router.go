package v1

import (
	"github.com/azoma13/computer-assembly-service/order/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(handler *echo.Echo, service *service.Services) {
	handler.Use(middleware.Recover())

	handler.GET("/health", func(c echo.Context) error { return c.NoContent(204) })

	v1 := handler.Group("/api/v1")
	{
		OrderRoutes(v1.Group("/order"), service.Order)
	}
}

type orderRoutes struct {
	orderService service.Order
}

func OrderRoutes(g *echo.Group, serviceOrder service.Order) {
	r := orderRoutes{
		orderService: serviceOrder,
	}

	g.POST("/create", r.createOrder)
}

func (r *orderRoutes) createOrder(c echo.Context) error {
	return nil
}
