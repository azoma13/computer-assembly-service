package pgdb

import (
	"sync"

	"github.com/azoma13/computer-assembly-service/shared/pkg/postgres"
)

type OrderRepo struct {
	mu       sync.RWMutex
	Postgres *postgres.Postgres
}

func NewOrderRepo(pg *postgres.Postgres) *OrderRepo {
	return &OrderRepo{Postgres: pg}
}
