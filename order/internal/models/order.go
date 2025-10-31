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
		UUID          string      `json:"uuid"`
		UserUUID      string      `json:"user_uuid"`
		HardwareUUIDs []string    `json:"hardware_uuids"`
		TotalPrice    float64     `json:"total_price"`
		Payment       *Payment    `json:"payment,omitempty"`
		Status        OrderStatus `json:"status"`
		UpdatedAt     *time.Time  `json:"updated_at,omitempty"`
		CreatedAt     time.Time   `json:"created_at"`
	}

	Payment struct {
		TransactionUUID string         `json:"transaction_uuid"`
		PaymentMethod   *PaymentMethod `json:"payment_method"`
		PaymentAt       time.Time      `json:"payment_at"`
	}

	PayOrderInfo struct {
		OrderUUID     string
		PaymentMethod *PaymentMethod
	}

	CreateOrderInfo struct {
		UserUUID      string
		HardwareUUIDs []string
	}

	CreateOrderData struct {
		OrderUUID  string
		TotalPrice float64
		Status     OrderStatus
		Created_at *time.Time
	}

	UpdateOrderInfo struct {
		TotalPrice      *float64
		TransactionUUID *string
		PaymentMethod   *PaymentMethod
		Status          *OrderStatus
	}
)
