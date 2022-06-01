package route

import (
	"github.com/labstack/echo/v4"
	_handlers "github.com/wildanie12/go-microservice/product-service/deliveries/handlers"
)

func RegisterProductRoutes(e *echo.Echo, productHandler *_handlers.ProductHandler) {
	e.POST("/products", productHandler.Create)
	e.GET("/products", productHandler.Index)
}