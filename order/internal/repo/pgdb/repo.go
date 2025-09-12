package pgdb

import (
	"sync"

	"github.com/azoma13/computer-assembly-service/order/pkg/postgres"
)

type OrderRepo struct {
	mu sync.RWMutex
	pg *postgres.Postgres
}

func NewOrderRepo(pg *postgres.Postgres) *OrderRepo {
	return &OrderRepo{
		pg: pg,
	}
}
