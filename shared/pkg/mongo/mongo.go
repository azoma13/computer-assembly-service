package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultMaxPoolSize  = 20
	defaultConnAttempts = 10
	defaultConnTimeout  = time.Second
)

type Mongo struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	MongoClient *mongo.Client
}

func New(url string, opts ...Option) (*Mongo, error) {
	mg := &Mongo{
		maxPoolSize:  defaultMaxPoolSize,
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	for _, option := range opts {
		option(mg)
	}

	clientOptions := options.Client().ApplyURI(url)
	clientOptions = clientOptions.SetMaxPoolSize(uint64(mg.maxPoolSize))

	var err error = nil
	for ; mg.connAttempts > 0; mg.connAttempts-- {
		mg.MongoClient, err = mongo.Connect(context.Background(), clientOptions)
		if err == nil {
			break
		}

		log.Printf("MongoDB is trying to connect, attempts left: %d", mg.connAttempts)
		time.Sleep(mg.connTimeout)
	}

	if err != nil {
		return nil, fmt.Errorf("error create connect mongodb: %w", err)
	}

	if err := mg.MongoClient.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return mg, nil
}

func (m *Mongo) Close() {
	if mongoErr := m.MongoClient.Disconnect(context.Background()); mongoErr != nil {
		log.Println("error: failed to disconnect from MongoDB: %w\n", mongoErr)
	} else {
		log.Println("Disconnect from MongoDB")
	}
}
