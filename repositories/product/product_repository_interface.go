package product

import _domain "github.com/wildanie12/go-microservice/product-service/entities/domain"

type ProductRepositoryInterface interface {

	// FindAll returns product resources in an array of entity struct
	// based on a given filter and sort query parameter 
	FindAll() ([]_domain.Product, error)

	// Insert will insert the given product parameter to data source
	// an error is returned when there was an error
	Insert(product _domain.Product) (_domain.Product, error)
}