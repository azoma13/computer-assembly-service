package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/azoma13/computer-assembly-service/order/internal/api/converter"
	"github.com/azoma13/computer-assembly-service/order/internal/models"
	order_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/openapi/order/v1"
)

func (a *orderAPI) PaymentOrder(ctx context.Context, req *order_v1.PaymentOrderRequest, params order_v1.PaymentOrderParams) (order_v1.PaymentOrderRes, error) {
	transactionUUID, err := a.orderService.PaymentOrder(ctx, params.OrderUUID.String(), string(req.GetPaymentMethod()))
	if err != nil {
		switch {
		case errors.Is(err, models.ErrPaymentOrderNotFound):
			return &order_v1.NotFoundError{
				Code:    http.StatusBadRequest,
				Message: "Заказ не найден или не существует",
			}, nil
		case errors.Is(err, models.ErrPaymentOrderConflict):
			return &order_v1.ConflictError{
				Code:    http.StatusBadRequest,
				Message: "Заказ уже оплачен или отменен",
			}, nil
		default:
			return &order_v1.InternalServerError{
				Code:    http.StatusInternalServerError,
				Message: "Внутренняя ошибка сервера",
			}, nil
		}
	}
	return &order_v1.PaymentOrderResponse{
		TransactionUUID: converter.StringToUUID(transactionUUID),
	}, nil
}
