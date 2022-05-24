package utils

import (
	"context"
	"fmt"

	_config "github.com/wildanie12/go-microservice/product-service/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoConnection(config *_config.Config) *mongo.Client {
	dsn := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		config.Mongo.User,
		config.Mongo.Pass,
		config.Mongo.Host,
	)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn))
	if err != nil {
		panic("Cannot connect to mongodb database")
	}

	return client
}