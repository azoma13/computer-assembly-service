package service

import (
	"context"

	"github.com/azoma13/computer-assembly-service/hardware/internal/models"
	"github.com/azoma13/computer-assembly-service/hardware/internal/repo"
	serviceHardware "github.com/azoma13/computer-assembly-service/hardware/internal/service/hardware"
)

type Hardware interface {
	GetHardware(ctx context.Context, hardwareUUID string) (models.Hardware, error)
	ListHardwares(ctx context.Context, filter models.HardwareFilter) ([]models.Hardware, error)
}

type Services struct {
	Hardware Hardware
}

type ServicesDependencies struct {
	Repos *repo.Repositories
}

func NewService(deps ServicesDependencies) *Services {
	return &Services{
		Hardware: serviceHardware.NewHardwareService(deps.Repos.Hardware),
	}
}
