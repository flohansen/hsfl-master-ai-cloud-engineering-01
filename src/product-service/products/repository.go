package products

import "github.com/flohansen/hsfl-master-ai-cloud-engineering/product-service/products/model"

type Repository interface {
	Create([]*model.Product) error
	FindAll() ([]*model.Product, error)
	FindById(id int64) (*model.Product, error)
	Delete([]*model.Product) error
}
