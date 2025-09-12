package service

import (
	"github.com/azoma13/computer-assembly-service/order/internal/repo"
)

type OrderService struct {
	orderRepo repo.OrderRepository
}

func NewOrderService(orderRepo repo.OrderRepository) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}
