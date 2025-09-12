package pgdb

import (
	"context"

	"github.com/azoma13/computer-assembly-service/order/internal/model"
	"github.com/jackc/pgx/v5"
)

func (r *OrderRepo) GetOrder(ctx context.Context, orderUUID string) (model.OrderData, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	query, args, err := r.pg.Builder.Select(
		"uuid",
		"user_uuid",
		"hardware_uuids",
		"total_price",
		"payment",
		"status",
		"updated_at",
		"created_at").
		From("orders").
		Where("uuid = ?", orderUUID).
		ToSql()
	if err != nil {
		return model.OrderData{}, err
	}

	var orderData model.OrderData
	err = r.pg.Pool.QueryRow(ctx, query, args...).Scan(
		orderData.UUID,
		orderData.UserUUID,
		orderData.HardwareUUIDs,
		orderData.TotalPrice,
		orderData.Payment,
		orderData.Status,
		orderData.UpdatedAt,
		orderData.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.OrderData{}, model.ErrOrderNotFound
		}
		return model.OrderData{}, err
	}

	return orderData, nil
}
