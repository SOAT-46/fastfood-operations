package configuration

import (
	"context"
	"time"

	"github.com/SOAT-46/fastfood-operations/internal/shared/domain/entities"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultTimeout = 10
)

func MongoClient(settings *entities.DatabaseSettings) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout*time.Second)
	defer cancel()

	mongoOptions := options.Client().ApplyURI(settings.GetMongoURI())
	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	return client.Database("fastfood_operations_database")
}
