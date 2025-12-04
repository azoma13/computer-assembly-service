package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/azoma13/computer-assembly-service/hardware/internal/models"
	"github.com/azoma13/computer-assembly-service/hardware/internal/repo/converter"
	repoModels "github.com/azoma13/computer-assembly-service/hardware/internal/repo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *hardwareRepo) GetHardware(ctx context.Context, hardwareUUID string) (models.Hardware, error) {
	var hardware repoModels.Hardware

	filter := bson.D{{Key: "uuid", Value: hardwareUUID}}

	err := r.collection.FindOne(context.Background(), filter).Decode(&hardware)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Hardware{}, models.ErrNotFoundHardware
		}
		return models.Hardware{}, fmt.Errorf("error findOne in collection hardware: %w", err)
	}

	return converter.HardwareToModel(hardware), nil
}
