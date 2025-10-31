package grpc

import (
	"context"

	"github.com/azoma13/computer-assembly-service/order/internal/models"
)

type HardwareClient interface {
	// GetHardware(ctx context.Context, hardwareUUID string) (models.Hardware, error)
	ListHardwares(ctx context.Context, filter models.HardwareFilter) ([]models.Hardware, error)
}

type PaymentClient interface {
	PayOrder(ctx context.Context, orderUUID, userUUID, paymentMethod string) (string, error)
}
