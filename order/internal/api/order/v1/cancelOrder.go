package v1

import (
	"context"
	"errors"

	"github.com/azoma13/computer-assembly-service/order/internal/models"
	order_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/openapi/order/v1"
)

func (a *orderAPI) CancelOrder(ctx context.Context, params order_v1.CancelOrderParams) (order_v1.CancelOrderRes, error) {
	err := a.orderService.CancelOrder(ctx, params.OrderUUID.String())
	if err != nil {
		switch {
		case errors.Is(err, models.ErrCancelOrderNotFound):
			return &order_v1.NotFoundError{
				Code:    404,
				Message: "Order by this UUID `" + params.OrderUUID.String() + "` not found",
			}, nil
		case errors.Is(err, models.ErrCancelOrderAlreadyPaid):
			return &order_v1.ConflictError{
				Code:    409,
				Message: "Заказ уже оплачен и не может быть отменён",
			}, nil
		case errors.Is(err, models.ErrCancelOrderAlreadyInProgress):
			return &order_v1.ConflictError{
				Code:    409,
				Message: "Заказ уже выполняется и не может быть отменён",
			}, nil
		case errors.Is(err, models.ErrCancelOrderAlreadyCompleted):
			return &order_v1.ConflictError{
				Code:    409,
				Message: "Заказ уже выполнен и не может быть отменён",
			}, nil
		case errors.Is(err, models.ErrCancelOrderAlreadyCancelled):
			return &order_v1.ConflictError{
				Code:    409,
				Message: "Заказ уже отменён",
			}, nil
		default:
			return &order_v1.InternalServerError{
				Code:    500,
				Message: "Внутренняя ошибка сервера",
			}, nil
		}
	}
	return nil, nil
}
