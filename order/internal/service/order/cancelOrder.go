package service

import (
	"context"

	"github.com/azoma13/computer-assembly-service/order/internal/models"
)

func (s *OrderService) CancelOrder(ctx context.Context, orderUUID string) error {
	order, err := s.orderRepo.GetOrder(ctx, orderUUID)
	if err != nil {
		return err
	}

	switch order.Status {
	case models.OrderStatusExpectPayment:
		status := models.OrderStatusCancelled
		err = s.orderRepo.UpdateOrder(ctx, orderUUID, models.UpdateOrderInfo{
			Status: &status,
		})
		if err != nil {
			return err
		}
		return nil
	case models.OrderStatusPaid:
		return models.ErrCancelOrderAlreadyPaid
	case models.OrderStatusInProgress:
		return models.ErrCancelOrderAlreadyInProgress
	case models.OrderStatusCompleted:
		return models.ErrCancelOrderAlreadyCompleted
	case models.OrderStatusCancelled:
		return models.ErrCancelOrderAlreadyCancelled
	default:
		return models.ErrCancelOrderInternalService
	}
}
