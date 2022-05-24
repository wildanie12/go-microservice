package main

import (
	"context"

	_config "github.com/wildanie12/go-microservice/product-service/config"
	_utils "github.com/wildanie12/go-microservice/product-service/utils"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	config := _config.New()


	client := _utils.NewMongoConnection(config)

	err := client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(err)
	}

	// e := echo.New()
	// e.Logger.Fatal(e.Start(":" + config.App.Port))
}