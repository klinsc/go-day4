package students

import (
	"github.com/klinsc/model"
	"gorm.io/gorm"
)

type IRepository interface {
	Create(student *model.Students) error
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &Repository{DB: db}
}

func (repo Repository) Create(student *model.Students) error {
	result := repo.DB.Create(student)

	return result.Error
}
