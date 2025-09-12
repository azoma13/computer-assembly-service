package pgdb

import (
	"context"
	"time"

	"github.com/azoma13/computer-assembly-service/order/internal/model"
	"github.com/azoma13/computer-assembly-service/order/internal/repo/converter"
	repoModel "github.com/azoma13/computer-assembly-service/order/internal/repo/model"
)

func (r *OrderRepo) CreateOrder(ctx context.Context, userUUID string, hardwares []model.Hardware) (model.CreateOrderData, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	hardwareUUIDs := make([]string, 0, len(hardwares))
	totalPrice := 0.00
	for _, hardware := range hardwares {
		hardwareUUIDs = append(hardwareUUIDs, hardware.UUID)
		totalPrice += hardware.Price
	}

	order := repoModel.OrderData{
		UserUUID:      userUUID,
		HardwareUUIDs: hardwareUUIDs,
		TotalPrice:    totalPrice,
		Status:        repoModel.OrderStatusExpectPayment,
		CreatedAt:     time.Now(),
	}

	query, args, err := r.pg.Builder.Insert("orders").
		Columns("user_uuid", "hardware_uuids", "total_price", "status", "created_at").
		Values(order.UserUUID, order.HardwareUUIDs, order.TotalPrice, order.Status, order.CreatedAt).
		Suffix("RETURN uuid, total_price, status, created_at").
		ToSql()
	if err != nil {
		return model.CreateOrderData{}, err
	}

	var info repoModel.CreateOrderData
	err = r.pg.Pool.QueryRow(ctx, query, args...).Scan(&info.OrderUUID, &info.TotalPrice, &info.Status, &info.Created_at)
	if err != nil {
		return model.CreateOrderData{}, err
	}

	return converter.CreateOrderInfoToModel(info), nil
}
