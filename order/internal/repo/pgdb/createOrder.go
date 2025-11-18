package pgdb

import (
	"context"
	"time"

	"github.com/azoma13/computer-assembly-service/order/internal/models"
	"github.com/azoma13/computer-assembly-service/order/internal/repo/converter"
	repoModels "github.com/azoma13/computer-assembly-service/order/internal/repo/models"
)

func (r *OrderRepo) CreateOrder(ctx context.Context, input models.CreateOrderInfo) (models.CreateOrderData, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var hardwareUUIDs []string
	totalPrice := 0.0
	for _, hardware := range input.HardwaresList {
		hardwareUUIDs = append(hardwareUUIDs, hardware.UUID)
		totalPrice += hardware.Price
	}

	order := repoModels.Order{
		UserUUID:      input.UserUUID,
		HardwareUUIDs: hardwareUUIDs,
		TotalPrice:    totalPrice,
		Status:        repoModels.OrderStatusExpectPayment,
		CreatedAt:     time.Now(),
	}

	result := r.Postgres.GormDB.Create(&order)
	if result.Error != nil {
		return models.CreateOrderData{}, nil
	}

	return converter.OrderCreateDataToModel(order), nil
}
