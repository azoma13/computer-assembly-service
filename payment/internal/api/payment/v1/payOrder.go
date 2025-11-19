package v1

import (
	"context"

	payment_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/payment/v1"
)

func (a *paymentAPI) PayOrder(ctx context.Context, req *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error) {
	transactionUUID, err := a.paymentService.PayOrder(ctx, req.GetOrderUuid())
	if err != nil {
		return nil, err
	}

	return &payment_v1.PayOrderResponse{
		TransactionUuid: transactionUUID,
	}, nil
}
