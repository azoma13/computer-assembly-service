package converter

import (
	"github.com/azoma13/computer-assembly-service/order/internal/models"
	repoModel "github.com/azoma13/computer-assembly-service/order/internal/repo/models"
)

func OrderGetDataToModel(order repoModel.Order) models.Order {
	modelPayment := models.Payment{
		TransactionUUID: order.Payment.TransactionUUID,
		PaymentMethod:   (*models.PaymentMethod)(order.Payment.PaymentMethod),
		PaymentAt:       order.Payment.PaymentAt,
	}

	return models.Order{
		UUID:          order.UUID,
		UserUUID:      order.UserUUID,
		HardwareUUIDs: order.HardwareUUIDs,
		TotalPrice:    order.TotalPrice,
		Payment:       &modelPayment,
		Status:        models.OrderStatus(order.Status),
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
	}
}

func OrderCreateDataToModel(order repoModel.Order) models.CreateOrderData {
	return models.CreateOrderData{
		OrderUUID:  order.UUID,
		TotalPrice: order.TotalPrice,
		Status:     models.OrderStatus(order.Status),
		Created_at: &order.CreatedAt,
	}
}
