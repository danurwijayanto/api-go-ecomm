package repository

import (
	"../entity"
)

// ProductRepository : Interface for product like CRUD
type ProductRepository interface {
	Add(product *entity.Product) (int, error)
	GetAll() ([]*entity.Product, error)
}
