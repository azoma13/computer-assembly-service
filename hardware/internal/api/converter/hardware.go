package converter

import (
	"github.com/azoma13/computer-assembly-service/hardware/internal/models"
	hardware_v1 "github.com/azoma13/computer-assembly-service/shared/pkg/proto/hardware/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func HardwareModelToProto(hardware models.Hardware) *hardware_v1.Hardware {
	var updateAt *timestamppb.Timestamp
	if hardware.UpdatedAt != nil {
		updateAt = timestamppb.New(*hardware.UpdatedAt)
	}

	return &hardware_v1.Hardware{
		Uuid:          hardware.UUID,
		Name:          hardware.Name,
		Description:   hardware.Description,
		Price:         hardware.Price,
		QuantityStock: hardware.QuantityStock,
		Category:      CategoryModelToProto(hardware.Category),
		Dimensions:    DimensionsModelToProto(hardware.Dimensions),
		Manufacturer:  ManufacturerModelToProto(hardware.Manufacturer),
		Metadata:      MetadataModelToProto(hardware.Metadata),
		UpdatedAt:     updateAt,
		CreatedAt:     timestamppb.New(hardware.CreatedAt),
	}
}

func ManyHardwareModelsToProto(hardwares []models.Hardware) []*hardware_v1.Hardware {
	res := make([]*hardware_v1.Hardware, 0, len(hardwares))
	for _, hardware := range hardwares {
		res = append(res, HardwareModelToProto(hardware))
	}

	return res
}

func MetadataModelToProto(metadata map[string]models.Metadata) map[string]*hardware_v1.Value {
	m := make(map[string]*hardware_v1.Value)
	for key, value := range metadata {
		switch {
		case value.StringValue != nil:
			m[key] = &hardware_v1.Value{
				Kind: &hardware_v1.Value_StringValue{StringValue: *value.StringValue},
			}
		case value.Int64Value != nil:
			m[key] = &hardware_v1.Value{
				Kind: &hardware_v1.Value_Int64Value{Int64Value: *value.Int64Value},
			}
		case value.DoubleValue != nil:
			m[key] = &hardware_v1.Value{
				Kind: &hardware_v1.Value_DoubleValue{DoubleValue: *value.DoubleValue},
			}
		case value.BoolValue != nil:
			m[key] = &hardware_v1.Value{
				Kind: &hardware_v1.Value_BoolValue{BoolValue: *value.BoolValue},
			}
		default:
			m[key] = &hardware_v1.Value{}
		}
	}
	return m
}

func CategoryModelToProto(category models.Category) hardware_v1.Category {
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

func CategoryProtoToModel(category hardware_v1.Category) models.Category {
	switch category {
	case hardware_v1.Category_CATEGORY_MOTHERBOARD:
		return models.CategoryMotherboard
	case hardware_v1.Category_CATEGORY_CPU:
		return models.CategoryCpu
	case hardware_v1.Category_CATEGORY_CPU_COOLER:
		return models.CategoryCpuCooler
	case hardware_v1.Category_CATEGORY_PSU:
		return models.CategoryPsu
	case hardware_v1.Category_CATEGORY_SSD:
		return models.CategorySsd
	case hardware_v1.Category_CATEGORY_HHD:
		return models.CategoryHhd
	case hardware_v1.Category_CATEGORY_RAM:
		return models.CategoryRam
	case hardware_v1.Category_CATEGORY_VIDEO_CARD:
		return models.CategoryVideoCard
	case hardware_v1.Category_CATEGORY_CASE:
		return models.CategoryCase
	default:
		return models.CategoryUnspecified
	}
}

func ManyCategoryProtoToModel(categories []hardware_v1.Category) []models.Category {
	res := make([]models.Category, 0, len(categories))
	for _, category := range categories {
		res = append(res, CategoryProtoToModel(category))
	}

	return res
}

func DimensionsModelToProto(dimensions models.Dimensions) *hardware_v1.Dimensions {
	return &hardware_v1.Dimensions{
		Length: dimensions.Length,
		Width:  dimensions.Width,
		Height: dimensions.Height,
		Weight: dimensions.Weight,
	}
}

func ManufacturerModelToProto(manufacturer models.Manufacturer) *hardware_v1.Manufacturer {
	return &hardware_v1.Manufacturer{
		Name:    manufacturer.Name,
		Country: manufacturer.Country,
		Website: manufacturer.Website,
	}
}

func HardwareFilterProtoToModel(filter *hardware_v1.HardwaresFilter) models.HardwareFilter {
	return models.HardwareFilter{
		UUIDs:                 filter.Uuids,
		Names:                 filter.Names,
		MinPrice:              filter.MinPrice,
		MaxPrice:              filter.MaxPrice,
		QuantityStock:         filter.QuantityStock,
		Categories:            ManyCategoryProtoToModel(filter.Categories),
		ManufacturerCountries: filter.ManufacturerCountries,
		Tags:                  filter.Tags,
	}
}
