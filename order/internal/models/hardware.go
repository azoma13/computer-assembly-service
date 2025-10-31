package models

import "time"

type Category string

const (
	CategoryUnspecified Category = "UNKNOWN"
	CategoryMotherboard Category = "MOTHERBOARD"
	CategoryCpu         Category = "CPU"
	CategoryCpuCooler   Category = "CPU_COOLER"
	CategoryPsu         Category = "PSU"
	CategorySsd         Category = "SSD"
	CategoryHhd         Category = "HHD"
	CategoryRam         Category = "RAM"
	CategoryVideoCard   Category = "VIDEO_CARD"
	CategoryCase        Category = "CASE"
)

type (
	Hardware struct {
		UUID          string
		Name          string
		Description   string
		Price         float64
		QuantityStock int64
		Category      Category
		Dimensions    Dimensions
		Manufacturer  Manufacturer
		Tags          []string
		Metadata      Metadata
		UpdatedAt     *time.Time
		CreatedAt     time.Time
	}

	Dimensions struct {
		Length float64
		Width  float64
		Height float64
		Weight float64
	}

	Manufacturer struct {
		Name    string
		Country string
		Website string
	}

	Metadata struct {
		StringValue *string
		Int64Value  *int64
		DoubleValue *float64
		BoolValue   *bool
	}

	HardwareFilter struct {
		UUIDs                 []string
		Names                 []string
		MinPrice              float64
		MaxPrice              float64
		QuantityStock         int64
		Categories            []Category
		ManufacturerCountries []string
		Tags                  []string
	}
)
