package user

import "github.com/flohansen/hsfl-master-ai-cloud-engineering/user-service/user/model"

type Repository interface {
	Migrate() error
	Create([]*model.DbUser) error
	FindByEmail(email string) ([]*model.DbUser, error)
	Delete([]*model.DbUser) error
}
