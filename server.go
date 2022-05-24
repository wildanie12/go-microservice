package main

import (
	"github.com/labstack/echo/v4"
	"github.com/wildanie12/go-microservice/product-service/config"
)

func main() {

	config := config.New()


	e := echo.New()
	

	e.Logger.Fatal(e.Start(":" + config.App.Port))
}