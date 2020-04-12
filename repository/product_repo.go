package repository

import "github.com/danurwijayanto/api-go-ecomm/entity"

// ProductRepository : Interface for product like CRUD
type ProductRepository interface {
	Add(product *entity.Product) (int, error)
	GetAll() ([]*entity.Product, error)
}
