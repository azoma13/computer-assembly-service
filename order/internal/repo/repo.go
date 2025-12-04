package repo

import (
	"context"

	"github.com/azoma13/computer-assembly-service/order/internal/models"
	"github.com/azoma13/computer-assembly-service/order/internal/repo/pgdb"
	"github.com/azoma13/computer-assembly-service/shared/pkg/postgres"
)

type Order interface {
	CreateOrder(ctx context.Context, input models.CreateOrderInfo) (models.CreateOrderData, error)
	GetOrder(ctx context.Context, orderUUID string) (models.Order, error)
	UpdateOrder(ctx context.Context, orderUUID string, input models.UpdateOrderInfo) error
}

type Repositories struct {
	Order
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		Order: pgdb.NewOrderRepo(pg),
	}
}
