package converter

import (
	"github.com/azoma13/computer-assembly-service/order/internal/model"
	repoModel "github.com/azoma13/computer-assembly-service/order/internal/repo/model"
)

func CreateOrderInfoToModel(createOrderInfo repoModel.CreateOrderData) model.CreateOrderData {
	return model.CreateOrderData{
		OrderUUID:  createOrderInfo.OrderUUID,
		TotalPrice: createOrderInfo.TotalPrice,
		Status:     model.OrderStatus(createOrderInfo.Status),
		Created_at: createOrderInfo.Created_at,
	}
}
