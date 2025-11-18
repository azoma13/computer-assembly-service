package service

import (
	"context"

	"github.com/azoma13/computer-assembly-service/order/internal/models"
)

func (s *OrderService) GetOrder(ctx context.Context, orderUUID string) (models.Order, error) {
	order, err := s.orderRepo.GetOrder(ctx, orderUUID)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}
