package services

import _web "github.com/wildanie12/go-microservice/product-service/entities/web"

type ProductServiceInterface interface {
	
	// Create will create a product resource 
	// with a given request parameter and returns proper resource response
	Create(productRequest _web.ProductRequest) (_web.ProductResponse, error)
}