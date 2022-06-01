package main

import (
	"github.com/labstack/echo/v4"
	_config "github.com/wildanie12/go-microservice/product-service/config"
	"github.com/wildanie12/go-microservice/product-service/deliveries/handlers"
	_route "github.com/wildanie12/go-microservice/product-service/deliveries/route"
	_productRepository "github.com/wildanie12/go-microservice/product-service/repositories/product"
	_productService "github.com/wildanie12/go-microservice/product-service/services/product"
	_utils "github.com/wildanie12/go-microservice/product-service/utils"
)

func main() {

	config := _config.New()
	mongodb := _utils.NewMongoConnection(config)

	productRepository := _productRepository.NewMongo(mongodb)
	productService := _productService.New(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	e := echo.New()
	_route.RegisterProductRoutes(e, productHandler)

	e.Logger.Fatal(e.Start(":" + config.App.Port))
}