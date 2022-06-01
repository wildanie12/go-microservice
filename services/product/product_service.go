package services

import (
	"github.com/jinzhu/copier"
	_domain "github.com/wildanie12/go-microservice/product-service/entities/domain"
	_web "github.com/wildanie12/go-microservice/product-service/entities/web"
	_productRepository "github.com/wildanie12/go-microservice/product-service/repositories/product"
)

type ProductService struct {
	productRepo _productRepository.ProductRepositoryInterface
}

func New(productRepository _productRepository.ProductRepositoryInterface) *ProductService {
	return &ProductService{
		productRepo: productRepository,
	}
}

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