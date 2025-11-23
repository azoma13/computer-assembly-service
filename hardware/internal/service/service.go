package service

import (
	"github.com/azoma13/computer-assembly-service/hardware/internal/repo"
	serviceHardware "github.com/azoma13/computer-assembly-service/hardware/internal/service/hardware"
)

type Hardware interface {
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
