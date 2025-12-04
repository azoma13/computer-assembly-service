package mongodb

import (
	"context"
	"log"
	"sync"
	"time"

	mongoPkg "github.com/azoma13/computer-assembly-service/shared/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type hardwareRepo struct {
	mu         sync.Mutex
	collection *mongo.Collection
}

func NewHardwareRepo(mg *mongoPkg.Mongo, db string) *hardwareRepo {
	collection := mg.MongoClient.Database(db).Collection("hardwares")

	indexModels := []mongo.IndexModel{
		{
			Keys: bson.D{{
				Key:   "uuid",
				Value: 1,
			}},
			Options: options.Index().SetUnique(true),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		log.Printf("failed to create indexes: %v", err)
	}

	s := &hardwareRepo{
		collection: collection,
	}

	return s
}
