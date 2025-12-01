package service

import (
	"github.com/azoma13/computer-assembly-service/order/internal/client/grpc"
	"github.com/azoma13/computer-assembly-service/order/internal/repo"
)

type OrderService struct {
	orderRepo repo.Order

	hardwareClient grpc.HardwareClient
	paymentClient  grpc.PaymentClient
}

func NewOrderService(orderRepo repo.Order, hardwareClient grpc.HardwareClient, paymentClient grpc.PaymentClient) *OrderService {
	return &OrderService{
		orderRepo:      orderRepo,
		hardwareClient: hardwareClient,
		paymentClient:  paymentClient,
	}
}
