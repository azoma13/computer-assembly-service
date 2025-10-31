package v1

import (
	"github.com/azoma13/computer-assembly-service/order/internal/client/grpc"
	payment_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/payment/v1"
)

var _ grpc.PaymentClient = (*client)(nil)

type client struct {
	generatedClient payment_v1.PaymentServiceClient
}

func NewPaymentClient(generatedClient payment_v1.PaymentServiceClient) *client {
	return &client{
		generatedClient: generatedClient,
	}
}
