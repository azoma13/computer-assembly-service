package converter

import (
	"time"

	"github.com/azoma13/computer-assembly-service/order/internal/models"
	hardware_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/hardware/v1"
	"github.com/samber/lo"
)

func ListHardwaresToModel(protoListHardwares []*hardware_v1.Hardware) []models.Hardware {
	var modelListHardwares []models.Hardware
	for _, hardware := range protoListHardwares {
		modelHardware := HardwareToModel(hardware)
		modelListHardwares = append(modelListHardwares, modelHardware)
	}

	return modelListHardwares
}

func HardwareToModel(protoHardware *hardware_v1.Hardware) models.Hardware {
	var updatedAt *time.Time
	if protoHardware != nil {
		updatedAt = lo.ToPtr(protoHardware.UpdatedAt.AsTime())
	}

	return models.Hardware{
		UUID:          protoHardware.Uuid,
		Name:          protoHardware.Name,
		Description:   protoHardware.Description,
		Price:         protoHardware.Price,
		QuantityStock: protoHardware.QuantityStock,
		Category:      models.Category(protoHardware.Category),
		Dimensions:    DimensionsToModel(protoHardware.Dimensions),
		Manufacturer:  ManufacturerToModel(protoHardware.Manufacturer),
		Tags:          protoHardware.Tags,
		Metadata:      MetadataToModel(protoHardware.Metadata),
		UpdatedAt:     updatedAt,
		CreatedAt:     protoHardware.CreatedAt.AsTime(),
	}
}

func DimensionsToModel(protoDimensions *hardware_v1.Dimensions) models.Dimensions {
	if protoDimensions == nil {
		return models.Dimensions{}
	}

	return models.Dimensions{
		Length: protoDimensions.Length,
		Width:  protoDimensions.Width,
		Height: protoDimensions.Height,
		Weight: protoDimensions.Weight,
	}
}

func ManufacturerToModel(protoManufacturer *hardware_v1.Manufacturer) models.Manufacturer {
	if protoManufacturer == nil {
		return models.Manufacturer{}
	}

	return models.Manufacturer{
		Name:    protoManufacturer.Name,
		Country: protoManufacturer.Country,
		Website: protoManufacturer.Website,
	}
}

func MetadataToModel(protoMetadata map[string]*hardware_v1.Value) models.Metadata {
	metadata := models.Metadata{}

	for _, value := range protoMetadata {
		if value == nil {
			continue
		}

		switch v := value.Kind.(type) {
		case *hardware_v1.Value_StringValue:
			metadata.StringValue = lo.ToPtr(v.StringValue)
		case *hardware_v1.Value_Int64Value:
			metadata.Int64Value = lo.ToPtr(v.Int64Value)
		case *hardware_v1.Value_BoolValue:
			metadata.BoolValue = lo.ToPtr(v.BoolValue)
		case *hardware_v1.Value_DoubleValue:
			metadata.DoubleValue = lo.ToPtr(v.DoubleValue)
		}
	}

	return metadata
}

func HardwareFilterToProto(modelHardwareFilter models.HardwareFilter) *hardware_v1.HardwaresFilter {
	return &hardware_v1.HardwaresFilter{
		Uuids:                 modelHardwareFilter.UUIDs,
		Names:                 modelHardwareFilter.Names,
		MinPrice:              modelHardwareFilter.MinPrice,
		MaxPrice:              modelHardwareFilter.MaxPrice,
		QuantityStock:         modelHardwareFilter.QuantityStock,
		Categories:            CategoriesToProto(modelHardwareFilter.Categories),
		ManufacturerCountries: modelHardwareFilter.ManufacturerCountries,
		Tags:                  modelHardwareFilter.Tags,
	}
}

func CategoriesToProto(protoCategories []models.Category) []hardware_v1.Category {
	if protoCategories == nil {
		return []hardware_v1.Category{}
	}

	res := []hardware_v1.Category{}
	for _, category := range protoCategories {
		res = append(res, CategoryToProto(category))
	}

	return res
}

func CategoryToProto(category models.Category) hardware_v1.Category {
	switch category {
	case models.CategoryMotherboard:
		return hardware_v1.Category_CATEGORY_MOTHERBOARD
	case models.CategoryCpu:
		return hardware_v1.Category_CATEGORY_CPU
	case models.CategoryCpuCooler:
		return hardware_v1.Category_CATEGORY_CPU_COOLER
	case models.CategoryPsu:
		return hardware_v1.Category_CATEGORY_PSU
	case models.CategorySsd:
		return hardware_v1.Category_CATEGORY_SSD
	case models.CategoryHhd:
		return hardware_v1.Category_CATEGORY_HHD
	case models.CategoryRam:
		return hardware_v1.Category_CATEGORY_RAM
	case models.CategoryVideoCard:
		return hardware_v1.Category_CATEGORY_VIDEO_CARD
	case models.CategoryCase:
		return hardware_v1.Category_CATEGORY_CASE
	default:
		return hardware_v1.Category_CATEGORY_UNSPECIFIED
	}
}
