package service

import (
	"github.com/azoma13/computer-assembly-service/order/internal/repo"
)

type Order interface {
	// CreateOrder(ctx context.Context, userUUID string, partsUUIDs []string) (info model.CreateOrderInfo, err error)
}

type Services struct {
	Order Order
}

type ServicesDependencies struct {
	Repos *repo.Repositories
}

func NewServices(deps ServicesDependencies) *Services {
	return &Services{
		Order: NewOrderService(deps.Repos.OrderRepository),
	}
}
