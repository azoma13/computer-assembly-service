package utils

import "github.com/azoma13/computer-assembly-service/order/internal/model"

func CheckStatusOrder(status model.OrderStatus) error {
	if status == model.OrderStatusPaid {
		return model.ErrOrderAlreadyPaid
	}

	if status == model.OrderStatusInProgress {
		return model.ErrOrderAlreadyInProgress
	}

	if status == model.OrderStatusCompleted {
		return model.ErrOrderAlreadyCompleted
	}

	if status == model.OrderStatusCancelled {
		return model.ErrOrderAlreadyCancelled
	}

	if status != model.OrderStatusExpectPayment {
		return model.ErrOrderUndefinedStatus
	}

	return nil
}
