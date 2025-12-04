package service

import (
	"context"

	"github.com/azoma13/computer-assembly-service/order/internal/models"
)

func (s *OrderService) CreateOrder(ctx context.Context, userUUID string, hardwareUUIDs []string) (models.CreateOrderData, error) {
	filter := models.HardwareFilter{
		UUIDs: hardwareUUIDs,
	}

	hardwaresList, err := s.hardwareClient.ListHardwares(ctx, filter)
	if err != nil {
		return models.CreateOrderData{}, models.ErrCreateOrderBadGateway
	}

	if len(hardwareUUIDs) != len(hardwaresList) {
		return models.CreateOrderData{}, models.ErrCreateOrderNotFound
	}

	orderData, err := s.orderRepo.CreateOrder(ctx, models.CreateOrderInfo{
		UserUUID:      userUUID,
		HardwaresList: hardwaresList,
	})
	if err != nil {
		return models.CreateOrderData{}, err
	}
	return orderData, nil
}
