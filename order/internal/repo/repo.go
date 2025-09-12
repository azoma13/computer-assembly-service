package repo

import (
	"context"

	"github.com/azoma13/computer-assembly-service/order/internal/model"
	"github.com/azoma13/computer-assembly-service/order/internal/repo/pgdb"

	"github.com/azoma13/computer-assembly-service/order/pkg/postgres"
)

var _ OrderRepository = (*pgdb.OrderRepo)(nil)

type OrderRepository interface {
	CreateOrder(ctx context.Context, userUUID string, hardwares []model.Hardware) (model.CreateOrderData, error)
	GetOrder(ctx context.Context, orderUUID string) (model.OrderData, error)
	UpdateOrder(ctx context.Context, orderUUID string, updateOrderInfo model.UpdateOrderInfo) error
}

type Repositories struct {
	OrderRepository
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		OrderRepository: pgdb.NewOrderRepo(pg),
	}
}
