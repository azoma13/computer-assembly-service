package v1

import (
	"github.com/azoma13/computer-assembly-service/order/internal/service"
)

type orderAPI struct {
	orderService service.Order
}

func NewOrderAPI(serviceOrder service.Order) *orderAPI {
	return &orderAPI{
		orderService: serviceOrder,
	}
}
