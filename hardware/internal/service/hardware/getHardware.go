package service

import (
	"context"

	"github.com/azoma13/computer-assembly-service/hardware/internal/models"
)

func (s *HardwareService) GetHardware(ctx context.Context, hardwareUUID string) (models.Hardware, error) {
	hardware, err := s.hardwareRepo.GetHardware(ctx, hardwareUUID)
	if err != nil {
		return models.Hardware{}, err
	}
	return hardware, nil
}
