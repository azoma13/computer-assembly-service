package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/azoma13/computer-assembly-service/order/internal/api/converter"
	"github.com/azoma13/computer-assembly-service/order/internal/models"
	order_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/openapi/order/v1"
)

func (a *orderAPI) CreateOrder(ctx context.Context, req *order_v1.CreateOrderRequest) (order_v1.CreateOrderRes, error) {
	orderInfo, err := a.orderService.CreateOrder(ctx, req.GetUserUUID().String(), converter.UUIDsToStrings(req.GetHardwareUuids()))
	if err != nil {
		switch {
		case errors.Is(err, models.ErrCreateOrderBadRequest):
			return &order_v1.BadRequestError{
				Code:    http.StatusBadRequest,
				Message: "Неверный запрос! Пожалуйста, попробуйте еще раз!",
			}, nil
		case errors.Is(err, models.ErrCreateOrderNotFound):
			return &order_v1.NotFoundError{
				Code:    http.StatusNotFound,
				Message: "Одна или несколько частей не найдены",
			}, nil
		case errors.Is(err, models.ErrCreateOrderBadGateway):
			return &order_v1.BadGatewayError{
				Code:    http.StatusBadGateway,
				Message: "Некорректный ответ от другого сервера",
			}, nil
		default:
			return &order_v1.InternalServerError{
				Code:    500,
				Message: "Внутренняя ошибка сервера",
			}, nil
		}
	}

	return &order_v1.CreateOrderResponse{
		OrderUUID:  converter.StringToUUID(orderInfo.OrderUUID),
		TotalPrice: orderInfo.TotalPrice,
	}, nil
}
