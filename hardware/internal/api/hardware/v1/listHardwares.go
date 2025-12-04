package v1

import (
	"context"

	"github.com/azoma13/computer-assembly-service/hardware/internal/api/converter"
	hardware_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/hardware/v1"
)

func (a *hardwareAPI) ListHardwares(ctx context.Context, req *hardware_v1.ListHardwaresRequest) (*hardware_v1.ListHardwaresResponse, error) {
	filter := converter.HardwareFilterProtoToModel(req.Filter)

	hardwares, err := a.hardwareService.ListHardwares(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &hardware_v1.ListHardwaresResponse{
		Hardwares: converter.ManyHardwareModelsToProto(hardwares),
	}, nil
}
