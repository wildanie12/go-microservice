package route

import (
	"github.com/labstack/echo/v4"
	_handlers "github.com/wildanie12/go-microservice/product-service/deliveries/handlers"
)

// RegisterProductRoutes makes a product route to `ProductHandler`
func RegisterProductRoutes(e *echo.Echo, productHandler *_handlers.ProductHandler) {
	e.POST("/products", productHandler.Create)
	e.GET("/products", productHandler.Index)
}