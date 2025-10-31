package repo

import (
	"github.com/azoma13/computer-assembly-service/order/internal/repo/pgdb"
	"github.com/azoma13/computer-assembly-service/order/pkg/postgres"
)

type Order interface {
}

type Repositories struct {
	Order
}

func NewRepositories(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		Order: pgdb.NewOrderRepo(pg),
	}
}
