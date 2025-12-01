package converter

import (
	"github.com/azoma13/computer-assembly-service/hardware/internal/models"
	repoModels "github.com/azoma13/computer-assembly-service/hardware/internal/repo/models"
)

func HardwareToModel(hardware repoModels.Hardware) models.Hardware {
	return models.Hardware{
		UUID:          hardware.UUID,
		Name:          hardware.Name,
		Description:   hardware.Description,
		Price:         hardware.Price,
		QuantityStock: hardware.QuantityStock,
		Category:      models.Category(hardware.Category),
		Dimensions:    HardwareDimensionsToModels(hardware.Dimensions),
		Manufacturer:  HardwareManufacturerToModels(hardware.Manufacturer),
		Tags:          hardware.Tags,
		Metadata:      HardwareMetadataToModels(hardware.Metadata),
		UpdatedAt:     hardware.UpdatedAt,
		CreatedAt:     hardware.CreatedAt,
	}
}

func HardwaresListToModels(hardwares []repoModels.Hardware) []models.Hardware {
	list := make([]models.Hardware, 0, len(hardwares))

	for _, hardware := range hardwares {
		list = append(list, HardwareToModel(hardware))
	}

	return list
}

func HardwareDimensionsToModels(dimensions repoModels.Dimensions) models.Dimensions {
	return models.Dimensions{
		Length: dimensions.Length,
		Width:  dimensions.Width,
		Height: dimensions.Height,
		Weight: dimensions.Weight,
	}
}

func HardwareManufacturerToModels(manufacturer repoModels.Manufacturer) models.Manufacturer {
	return models.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		Website: manufacturer.Website,
	}
}

func HardwareMetadataToModels(metadata map[string]repoModels.Metadata) map[string]models.Metadata {
	m := make(map[string]models.Metadata)

	for key, value := range metadata {
		m[key] = models.Metadata{
			StringValue: value.StringValue,
			Int64Value:  value.Int64Value,
			DoubleValue: value.DoubleValue,
			BoolValue:   value.BoolValue,
		}
	}

	return m
}
