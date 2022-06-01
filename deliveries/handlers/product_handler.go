package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_web "github.com/wildanie12/go-microservice/product-service/entities/web"
	_productService "github.com/wildanie12/go-microservice/product-service/services/product"
)

type ProductHandler struct {
	productService _productService.ProductServiceInterface
}

func NewProductHandler(productService _productService.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}


func (handler ProductHandler) Create(c echo.Context) error {

	productReq := _web.ProductRequest {
		Name: "Sades Lance Mouse",
		Unit: "pcs",
		Price: 28000,
		Category: "Computer Accessories",
	}

	productRes, err := handler.productService.Create(productReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{} {
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, productRes)
}
