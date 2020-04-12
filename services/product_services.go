package services

import (
	"github.com/danurwijayanto/api-go-ecomm/entity"
	"github.com/danurwijayanto/api-go-ecomm/repository"
)

// ProductServices :
type ProductServices interface {
	Add(product *entity.Product) (int, error)
	GetAll() ([]*entity.Product, error)
}

type productServices struct{}

var (
	productRepository repository.ProductRepository = repository.NewMysqlProductRepository()
)

// NewProductServices : Constructor
func NewProductServices() ProductServices {
	return &productServices{}
}

func (*productServices) Add(product *entity.Product) (int, error) {
	return productRepository.Add(product)
}

func (*productServices) GetAll() ([]*entity.Product, error) {
	return productRepository.GetAll()
}
