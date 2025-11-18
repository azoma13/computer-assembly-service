package pgdb

import (
	"context"
	"fmt"
	"time"

	"github.com/azoma13/computer-assembly-service/order/internal/models"
	repoModels "github.com/azoma13/computer-assembly-service/order/internal/repo/models"
)

func (r *OrderRepo) UpdateOrder(ctx context.Context, orderUUID string, input models.UpdateOrderInfo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	sel := make([]string, 0, 4)
	order := repoModels.Order{}
	if input.TotalPrice != nil {
		sel = append(sel, "total_price")
		order.TotalPrice = *input.TotalPrice
	}

	if input.TransactionUUID != nil {
		sel = append(sel, "payment")
		order.Payment.TransactionUUID = *input.TransactionUUID
		order.Payment.PaymentMethod = (*repoModels.PaymentMethod)(input.PaymentMethod)
		order.Payment.PaymentAt = time.Now()
	}

	if input.Status != nil {
		sel = append(sel, "status")
		order.Status = repoModels.OrderStatus(*input.Status)
	}

	result := r.Postgres.GormDB.Model(&repoModels.Order{}).Where("uuid = ?", orderUUID).Select(sel).Updates(order)
	if result.Error != nil {
		return fmt.Errorf("failed to update order: %w", result.Error)
	}
	return nil
}
