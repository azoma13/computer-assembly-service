package pgdb

import (
	"context"
	"errors"

	"github.com/azoma13/computer-assembly-service/order/internal/models"
	"github.com/azoma13/computer-assembly-service/order/internal/repo/converter"
	repoModels "github.com/azoma13/computer-assembly-service/order/internal/repo/models"
	"gorm.io/gorm"
)

func (r *OrderRepo) GetOrder(ctx context.Context, orderUUID string) (models.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var order repoModels.Order
	res := r.Postgres.GormDB.First(order, "uuid = ?", orderUUID)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return models.Order{}, models.ErrGetOrderNotFound
		}
		return models.Order{}, res.Error
	}

	if res.RowsAffected != 1 {
		return models.Order{}, errors.New("multiple orders found")
	}

	return converter.OrderGetDataToModel(order), nil
}
