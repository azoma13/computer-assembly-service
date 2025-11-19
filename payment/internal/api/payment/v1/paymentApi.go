package v1

import (
	"github.com/azoma13/computer-assembly-service/payment/internal/service"
	payment_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/payment/v1"
)

type paymentAPI struct {
	payment_v1.UnimplementedPaymentServiceServer
	paymentService service.Payment
}

func NewPaymentAPI(service service.Services) *paymentAPI {
	return &paymentAPI{
		paymentService: service.Payment,
	}
}
