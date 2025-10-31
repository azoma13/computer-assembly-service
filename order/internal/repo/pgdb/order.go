package pgdb

import "github.com/azoma13/computer-assembly-service/order/pkg/postgres"

type OrderRepo struct {
	*postgres.Postgres
}

func NewOrderRepo(pg *postgres.Postgres) *OrderRepo {
	return &OrderRepo{pg}
}
