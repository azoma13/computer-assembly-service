package repo

import (
	"github.com/azoma13/computer-assembly-service/hardware/internal/repo/mongodb"
	mongoPkg "github.com/azoma13/computer-assembly-service/shared/pkg/mongo"
)

type Hardware interface {
	// GetHardware(ctx context.Context, hardwareUUID string) (models.Hardware, error)
	// ListHardwares(ctx context.Context, filter models.HardwareFilter) ([]models.Hardware, error)
}

type Repositories struct {
	Hardware Hardware
}

func NewRepositories(mg *mongoPkg.Mongo) *Repositories {
	return &Repositories{
		Hardware: mongodb.NewHardwareRepo(mg),
	}
}
