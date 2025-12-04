package converter

import (
	"log"

	"github.com/azoma13/computer-assembly-service/order/internal/models"
	order_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/openapi/order/v1"
	"github.com/google/uuid"
)

func ModelOrderToResOrder(order models.Order) order_v1.Order {
	payment := order_v1.NewOptPayment(order_v1.Payment{
		TransactionUUID: StringToUUID(order.Payment.TransactionUUID),
		PaymentMethod:   order_v1.PaymentMethod(*order.Payment.PaymentMethod),
		PaymentAt:       order_v1.NewOptDateTime(order.Payment.PaymentAt),
	})

	return order_v1.Order{
		OrderUUID:     StringToUUID(order.UUID),
		UserUUID:      StringToUUID(order.UserUUID),
		HardwareUuids: StringsToUUIDs(order.HardwareUUIDs),
		TotalPrice:    order.TotalPrice,
		Payment:       payment,
		Status:        order_v1.Status(order.Status),
		UpdatedAt:     order_v1.NewOptDateTime(*order.UpdatedAt),
		CreatedAt:     order.CreatedAt,
	}
}

func StringToUUID(s string) uuid.UUID {
	uuid, err := uuid.Parse(s)
	if err != nil {
		log.Printf("Failed to parse UUID: %v", err)
	}

	return uuid
}

func StringsToUUIDs(arr []string) []uuid.UUID {
	uuids := make([]uuid.UUID, len(arr))
	for i, s := range arr {
		uuids[i] = StringToUUID(s)
	}
	return uuids
}

func UUIDsToStrings(arr []uuid.UUID) []string {
	uuids := make([]string, len(arr))
	for i, s := range arr {
		uuids[i] = s.String()
	}
	return uuids
}
