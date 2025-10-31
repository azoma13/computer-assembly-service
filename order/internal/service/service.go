package service

import (
	"github.com/azoma13/computer-assembly-service/order/internal/client/grpc"
	"github.com/azoma13/computer-assembly-service/order/internal/repo"
)

type Order interface{}

type Services struct {
	Order Order
}

type ServicesDependencies struct {
	Repos *repo.Repositories

	HardwareClient grpc.HardwareClient
	PaymentClient  grpc.PaymentClient
}

func NewService(deps ServicesDependencies) *Services {
	return &Services{
		Order: NewOrderService(deps.Repos.Order, deps.HardwareClient, deps.PaymentClient),
	}
}
