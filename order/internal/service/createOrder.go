package service

import (
	"context"

	"github.com/azoma13/computer-assembly-service/order/internal/model"
)

func (s *OrderService) CreateOrder(ctx context.Context, input model.CreateOrderInfo) (model.CreateOrderData, error) {

	// get hardwares from inventory
	// check conflict

	order, err := s.orderRepo.CreateOrder(ctx, input.UserUUID, []model.Hardware{})
	if err != nil {
		return order, err
	}

	return order, nil
}
