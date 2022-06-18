package services

import _web "github.com/wildanie12/go-microservice/product-service/entities/web"

// ServiceInterface provide contract for product service
type ServiceInterface interface {
	
	// FindAll returns an array of product responses based on
	// defined filter parameter, 
	FindAll(filters []map[string]string) ([]_web.ProductResponse, error)

	// Create will create a product resource 
	// with a given request parameter and returns proper resource response
	Create(productRequest _web.ProductRequest) (_web.ProductResponse, error)
}