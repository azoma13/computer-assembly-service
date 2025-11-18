package service

import (
	"context"

	"github.com/azoma13/computer-assembly-service/order/internal/client/grpc"
	"github.com/azoma13/computer-assembly-service/order/internal/models"
	"github.com/azoma13/computer-assembly-service/order/internal/repo"
	serviceOrder "github.com/azoma13/computer-assembly-service/order/internal/service/order"
)

type Order interface {
	CreateOrder(ctx context.Context, userUUID string, hardwareUUIDs []string) (models.CreateOrderData, error)
	GetOrder(ctx context.Context, orderUUID string) (models.Order, error)
	CancelOrder(ctx context.Context, orderUUID string) error
	PaymentOrder(ctx context.Context, orderUUID, paymentMethod string) (string, error)
}

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
		Order: serviceOrder.NewOrderService(deps.Repos.Order, deps.HardwareClient, deps.PaymentClient),
	}
}
