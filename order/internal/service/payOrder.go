package service

import (
	"context"
	"fmt"

	"github.com/azoma13/computer-assembly-service/order/internal/model"
	"github.com/azoma13/computer-assembly-service/order/internal/utils"
)

func (s *OrderService) PayOrder(ctx context.Context, input model.PayOrderInfo) (string, error) {
	order, err := s.orderRepo.GetOrder(ctx, input.OrderUUID)
	if err != nil {
		return "", err
	}

	if err := utils.CheckStatusOrder(order.Status); err != nil {
		return "", fmt.Errorf("%s, cannot be paid", err)
	}

	// get transaction uuid from payment
	transUUID := ""

	status := model.OrderStatusPaid
	err = s.orderRepo.UpdateOrder(ctx, input.OrderUUID, model.UpdateOrderInfo{
		TransactionUUID: &transUUID, // заменить
		PaymentMethod:   input.PaymentMethod,
		Status:          &status,
	})
	if err != nil {
		return "", err
	}

	return "uuid", nil
}
