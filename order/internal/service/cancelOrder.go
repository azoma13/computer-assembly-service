package service

import (
	"context"
	"fmt"

	"github.com/azoma13/computer-assembly-service/order/internal/model"
	"github.com/azoma13/computer-assembly-service/order/internal/utils"
)

func (s *OrderService) CancelOrder(ctx context.Context, orderUUID string) error {
	order, err := s.orderRepo.GetOrder(ctx, orderUUID)
	if err != nil {
		return fmt.Errorf("CancelOrder - s.orderRepo.GetOrder: %v", err)
	}

	if err = utils.CheckStatusOrder(order.Status); err != nil {
		return fmt.Errorf("%s, cannot be cancelled", err)
	}

	status := model.OrderStatusCancelled
	err = s.orderRepo.UpdateOrder(ctx, orderUUID, model.UpdateOrderInfo{
		Status: &status,
	})
	if err != nil {
		return err
	}

	return nil
}
