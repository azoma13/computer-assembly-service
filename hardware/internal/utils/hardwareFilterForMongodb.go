package utils

import (
	"github.com/azoma13/computer-assembly-service/hardware/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func HardwareFilter(filter models.HardwareFilter) bson.M {
	m := bson.M{}

	if len(filter.UUIDs) != 0 {
		m["uuid"] = bson.M{"$in": filter.UUIDs}
	}

	if len(filter.Names) != 0 {
		m["name"] = bson.M{"$in": filter.Names}
	}

	if filter.MinPrice != 0 {
		m["price"] = bson.M{"$gte": filter.MinPrice}
	}

	if filter.MaxPrice != 0 {
		m["price"] = bson.M{"$lte": filter.MaxPrice}
	}

	if filter.QuantityStock != 0 {
		m["quantity_stock"] = bson.M{"$gte": filter.QuantityStock}
	}

	if len(filter.Categories) > 0 {
		m["category"] = bson.M{"$in": filter.Categories}
	}
	if len(filter.ManufacturerCountries) > 0 {
		m["manufacturer.country"] = bson.M{"$in": filter.ManufacturerCountries}
	}
	if len(filter.Tags) > 0 {
		m["tags"] = bson.M{"$all": filter.Tags}
	}

	return m
}
