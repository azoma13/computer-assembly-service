package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
		ID            primitive.ObjectID  `bson:"_id,omitempty"`
		UUID          string              `bson:"uuid"`
		Name          string              `bson:"name"`
		Description   string              `bson:"description"`
		Price         float64             `bson:"price"`
		QuantityStock int64               `bson:"quantity_stock"`
		Category      Category            `bson:"category"`
		Dimensions    Dimensions          `bson:"dimensions"`
		Manufacturer  Manufacturer        `bson:"manufacturer"`
		Tags          []string            `bson:"tags"`
		Metadata      map[string]Metadata `bson:"metadata"`
		UpdatedAt     *time.Time          `bson:"updatedAt"`
		CreatedAt     time.Time           `bson:"createdAt"`
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
