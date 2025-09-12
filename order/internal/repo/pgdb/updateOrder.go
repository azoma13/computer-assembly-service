package pgdb

import (
	"context"
	"time"

	"github.com/azoma13/computer-assembly-service/order/internal/model"
	"github.com/samber/lo"
)

func (r *OrderRepo) UpdateOrder(ctx context.Context, orderUUID string, updateOrderInfo model.UpdateOrderInfo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	now := time.Now()
	updateBuilder := r.pg.Builder.Update("orders").Set("update_at", now)

	if updateOrderInfo.TotalPrice != nil {
		updateBuilder = updateBuilder.Set("total_price", lo.FromPtr(updateOrderInfo.TotalPrice))
	}

	if updateOrderInfo.TransactionUUID != nil {
		method := model.PaymentMethodUndefined
		payment := &model.Payment{
			TransactionUUID: lo.FromPtr(updateOrderInfo.TransactionUUID),
			PaymentMethod:   lo.FromPtrOr(&updateOrderInfo.PaymentMethod, &method),
			PaymentAt:       now,
		}
		updateBuilder = updateBuilder.Set("payment", lo.FromPtr(payment))
		updateBuilder = updateBuilder.Set("status", lo.FromPtrOr(updateOrderInfo.Status, model.OrderStatusPaid))
	}

	if updateOrderInfo.Status != nil {
		updateBuilder = updateBuilder.Set("status", lo.FromPtrOr(updateOrderInfo.Status, model.OrderStatusExpectPayment))
	}

	updateBuilder.Where("uuid = ?", orderUUID)

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pg.Pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
