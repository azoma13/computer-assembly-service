package service

import "github.com/azoma13/computer-assembly-service/hardware/internal/repo"

type HardwareService struct {
	hardwareRepo repo.Hardware
}

func NewHardwareService(hardwareRepo repo.Hardware) *HardwareService {
	return &HardwareService{
		hardwareRepo: hardwareRepo,
	}
}
