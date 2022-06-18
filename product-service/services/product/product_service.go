package services

import (
	"github.com/jinzhu/copier"
	_domain "github.com/wildanie12/go-microservice/product-service/entities/domain"
	_web "github.com/wildanie12/go-microservice/product-service/entities/web"
	_productRepository "github.com/wildanie12/go-microservice/product-service/repositories/product"
)

// ProductService main struct
type ProductService struct {
	productRepo _productRepository.RepositoryInterface
}

// New acts as a ProductService constructor
func New(productRepository _productRepository.RepositoryInterface) *ProductService {
	return &ProductService{
		productRepo: productRepository,
	}
}

// FindAll returns an array of product responses based on
// defined filter parameter, 
func (service ProductService) FindAll(filters []map[string]string) ([]_web.ProductResponse, error) {
	products, err := service.productRepo.FindAll(filters)
	if err != nil {
		return []_web.ProductResponse{}, err
	}


	productsRes := []_web.ProductResponse{}
	copier.Copy(&productsRes, &products)
	return productsRes, nil
}

// Create method creates product resource with specified product request handler
// a series of validation and business logic can happen in this method
func (service ProductService) Create(productRequest _web.ProductRequest) (_web.ProductResponse, error) {

	product := _domain.Product{}

	copier.Copy(&product, &productRequest)
	product, err := service.productRepo.Insert(product)
	if err != nil {
		return _web.ProductResponse{}, err
	}

	productResponse := _web.ProductResponse{}
	copier.Copy(&productResponse, &product)
	return productResponse, nil
}