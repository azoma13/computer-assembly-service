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
		HardwareUUIDs []string
		TotalPrice    float64
		Payment       *Payment
		Status        OrderStatus
		UpdatedAt     *time.Time
		CreatedAt     time.Time
	}

	Payment struct {
		TransactionUUID string
		PaymentMethod   *PaymentMethod
		PaymentAt       time.Time
	}

	PayOrderInfo struct {
		OrderUUID     string
		PaymentMethod *PaymentMethod
	}

	CreateOrderInfo struct {
		UserUUID      string
		HardwaresList []Hardware
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
