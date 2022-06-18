package utils

import (
	"context"
	"fmt"

	_config "github.com/wildanie12/go-microservice/product-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoConnection return new connection instance `mongo.Database` to mongodb database
func NewMongoConnection(config *_config.Config) *mongo.Database {
	dsn := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?retryWrites=true&w=majority",
		config.Mongo.User,
		config.Mongo.Pass,
		config.Mongo.Host,
		config.Mongo.Port,
	)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn))
	if err != nil {
		panic("Cannot connect to mongodb database")
	}

	clientDB := client.Database(config.Mongo.DB)
	return clientDB
}