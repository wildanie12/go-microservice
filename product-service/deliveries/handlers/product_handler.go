package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_web "github.com/wildanie12/go-microservice/product-service/entities/web"
	_productService "github.com/wildanie12/go-microservice/product-service/services/product"
)

// ProductHandler main struct
type ProductHandler struct {
	productService _productService.ProductServiceInterface
}

// NewProductHandler act as a constructor 
// to return new instance of `ProductHandler`
func NewProductHandler(productService _productService.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// Index serves as list product endpoint controller implementation 
func (handler ProductHandler) Index(c echo.Context) error {
	products, err := handler.productService.FindAll([]map[string]string{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "OK",
		"code": http.StatusOK,
		"data": products,
	})
}

// Create responsible to create a product resource endpoint
// that can be hit using defined request body as a payload data
func (handler ProductHandler) Create(c echo.Context) error {

	productReq :=_web.ProductRequest{}
	c.Bind(&productReq)

	productRes, err := handler.productService.Create(productReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{} {
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, productRes)
}
