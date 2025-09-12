package service

import (
	"context"
	"fmt"

	"github.com/azoma13/computer-assembly-service/order/internal/model"
)

func (s *OrderService) GetOrder(ctx context.Context, orderUUID string) (model.OrderData, error) {
	order, err := s.orderRepo.GetOrder(ctx, orderUUID)
	if err != nil {
		return model.OrderData{}, fmt.Errorf("GetOrder - s.orderRepo.GetOrder: %v", err)
	}

	return order, nil
}
