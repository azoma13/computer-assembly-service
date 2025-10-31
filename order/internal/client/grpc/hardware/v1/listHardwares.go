package v1

import (
	"context"

	clientConverter "github.com/azoma13/computer-assembly-service/order/internal/client/converter"
	"github.com/azoma13/computer-assembly-service/order/internal/models"
	hardware_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/hardware/v1"
)

func (c *client) ListHardwares(ctx context.Context, filter models.HardwareFilter) ([]models.Hardware, error) {
	hardwaresFilter := &hardware_v1.ListHardwaresRequest{
		Filter: clientConverter.HardwareFilterToProto(filter),
	}

	listHardwares, err := c.generatedClient.ListHardwares(ctx, hardwaresFilter)
	if err != nil {
		return nil, err
	}

	return clientConverter.ListHardwaresToModel(listHardwares.Hardwares), nil
}
