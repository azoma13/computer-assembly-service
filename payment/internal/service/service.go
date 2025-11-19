package service

import (
	"context"

	servicePayment "github.com/azoma13/computer-assembly-service/payment/internal/service/payment"
)

type Payment interface {
	PayOrder(ctx context.Context, orderUUID string) (string, error)
}

type Services struct {
	Payment Payment
}

type ServicesDependencies struct {
}

func NewService(deps ServicesDependencies) *Services {
	return &Services{
		Payment: servicePayment.NewPaymentService(),
	}
}
