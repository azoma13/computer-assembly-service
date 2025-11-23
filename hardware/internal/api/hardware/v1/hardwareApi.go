package v1

import (
	"github.com/azoma13/computer-assembly-service/hardware/internal/service"
	hardware_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/hardware/v1"
)

type hardwareAPI struct {
	hardware_v1.UnimplementedHardwareServiceServer
	hardwareService service.Hardware
}

func NewHardwareAPI(service service.Services) *hardwareAPI {
	return &hardwareAPI{
		hardwareService: service.Hardware,
	}
}
