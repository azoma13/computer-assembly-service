package v1

import (
	"context"
	"errors"

	"github.com/azoma13/computer-assembly-service/order/internal/api/converter"
	"github.com/azoma13/computer-assembly-service/order/internal/models"
	order_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/openapi/order/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *orderAPI) GetOrder(ctx context.Context, params order_v1.GetOrderParams) (order_v1.GetOrderRes, error) {
	order, err := a.orderService.GetOrder(ctx, params.OrderUUID.String())
	if err != nil {
		switch {
		case errors.Is(err, models.ErrGetOrderNotFound):
			return nil, status.Errorf(codes.NotFound, "order by this UUID %s not found", params.OrderUUID.String())
		case errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled):
			return nil, status.Errorf(codes.Unavailable, "Order service timeout")
		default:
			return nil, status.Errorf(codes.Internal, "Order service internal error")
		}
	}
	return &order_v1.GetOrderResponse{
		Order: converter.ModelOrderToResOrder(order),
	}, nil
}
