package models

import "time"

type OrderStatus string

const (
	OrderStatusExpectPayment OrderStatus = "EXPECT_PAYMENT"
	OrderStatusPaid          OrderStatus = "PAID"
	OrderStatusInProgress    OrderStatus = "IN_PROGRESS"
	OrderStatusCompleted     OrderStatus = "COMPLETED"
	OrderStatusCancelled     OrderStatus = "CANCELLED"
)

type PaymentMethod string

const (
	PaymentMethodUndefined PaymentMethod = "UNDEFINED"
	PaymentMethodCash      PaymentMethod = "CASH"
	PaymentMethodCard      PaymentMethod = "CARD"
	PaymentMethodSbp       PaymentMethod = "SBP"
)

type (
	Order struct {
		UUID          string
		UserUUID      string
		HardwareUUIDs []string `gorm:"type:json"`
		TotalPrice    float64
		Payment       *Payment `gorm:"type:json"`
		Status        OrderStatus
		UpdatedAt     *time.Time
		CreatedAt     time.Time
	}

	Payment struct {
		TransactionUUID string         `json:"transaction_uuid"`
		PaymentMethod   *PaymentMethod `json:"payment_method"`
		PaymentAt       time.Time      `json:"payment_at"`
	}
)
