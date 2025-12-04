package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/azoma13/computer-assembly-service/hardware/internal/models"
	"github.com/azoma13/computer-assembly-service/hardware/internal/repo/converter"
	repoModels "github.com/azoma13/computer-assembly-service/hardware/internal/repo/models"
	"github.com/azoma13/computer-assembly-service/hardware/internal/utils"
)

func (r *hardwareRepo) ListHardwares(ctx context.Context, hardwareFilter models.HardwareFilter) ([]models.Hardware, error) {
	filter := utils.HardwareFilter(hardwareFilter)

	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer func() {
		cErr := cur.Close(ctx)
		if cErr != nil {
			log.Printf("failed to close cursor find hardwares with filter: %v\n", err)
		}
	}()

	res := []repoModels.Hardware{}

	err = cur.All(ctx, &res)
	if err != nil {
		return nil, fmt.Errorf("error all iterate cursor")
	}

	return converter.HardwaresListToModels(res), nil
}
