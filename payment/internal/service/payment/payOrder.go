package payment

import (
	"context"

	"github.com/google/uuid"
)

func (s *PaymentService) PayOrder(ctx context.Context, orderUUID string) (string, error) {
	uuid := uuid.New().String()

	return uuid, nil
}
