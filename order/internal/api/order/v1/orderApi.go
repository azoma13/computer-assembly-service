package v1

import (
	"github.com/azoma13/computer-assembly-service/order/internal/service"
)

type orderAPI struct {
	orderService service.Order
}

func NewOrderAPI(service service.Services) *orderAPI {
	return &orderAPI{
		orderService: service.Order,
	}
}
