package v1

import (
	"context"

	payment_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/payment/v1"
)

func (c *client) PayOrder(ctx context.Context, orderUUID, userUUID, paymentMethod string) (string, error) {
	input := &payment_v1.PayOrderRequest{
		OrderUuid:     orderUUID,
		UserUuid:      userUUID,
		PaymentMethod: payment_v1.PaymentMethod(payment_v1.PaymentMethod_value[paymentMethod]),
	}
	transactionUUID, err := c.generatedClient.PayOrder(ctx, input)
	if err != nil {
		return "", err
	}
	return transactionUUID.TransactionUuid, nil
}
