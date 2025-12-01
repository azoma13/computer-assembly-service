package v1

import (
	"context"
	"errors"

	"github.com/azoma13/computer-assembly-service/hardware/internal/api/converter"
	"github.com/azoma13/computer-assembly-service/hardware/internal/models"
	hardware_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/hardware/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *hardwareAPI) GetHardware(ctx context.Context, req *hardware_v1.GetHardwareRequest) (*hardware_v1.GetHardwareResponse, error) {
	hardware, err := a.hardwareService.GetHardware(ctx, req.Uuid)
	if err != nil {
		if errors.Is(err, models.ErrNotFoundHardware) {
			return nil, status.Errorf(codes.NotFound, "hardware not found: uuid{%s}", req.Uuid)
		}
		return nil, status.Errorf(codes.Internal, "internal error hardware server")
	}
	return &hardware_v1.GetHardwareResponse{
		Hardware: converter.HardwareModelToProto(hardware),
	}, nil
}
