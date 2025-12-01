package service

import (
	"context"

	"github.com/azoma13/computer-assembly-service/hardware/internal/models"
)

func (s *HardwareService) ListHardwares(ctx context.Context, hardwareFilter models.HardwareFilter) ([]models.Hardware, error) {
	hardwares, err := s.hardwareRepo.ListHardwares(ctx, hardwareFilter)
	if err != nil {
		return nil, err
	}
	return hardwares, nil
}
