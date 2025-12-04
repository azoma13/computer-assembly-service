package service

import (
	"context"

	"github.com/azoma13/computer-assembly-service/order/internal/models"
)

func (s *OrderService) PaymentOrder(ctx context.Context, orderUUID, paymentMethod string) (string, error) {
	order, err := s.orderRepo.GetOrder(ctx, orderUUID)
	if err != nil {
		return "", err
	}

	switch order.Status {
	case models.OrderStatusPaid:
		return "", models.ErrPaymentOrderConflict
	case models.OrderStatusInProgress:
		return "", models.ErrPaymentOrderConflict
	case models.OrderStatusCompleted:
		return "", models.ErrPaymentOrderConflict
	case models.OrderStatusCancelled:
		return "", models.ErrPaymentOrderConflict
	default:
	}

	transactionUUID, err := s.paymentClient.PayOrder(ctx, orderUUID, order.UserUUID, paymentMethod)
	if err != nil {
		return "", err
	}

	orderStatus := models.OrderStatusPaid
	err = s.orderRepo.UpdateOrder(ctx, orderUUID, models.UpdateOrderInfo{
		TransactionUUID: &transactionUUID,
		PaymentMethod:   (*models.PaymentMethod)(&paymentMethod),
		Status:          &orderStatus,
	})

	if err != nil {
		return "", err
	}

	return transactionUUID, nil
}
