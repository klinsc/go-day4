package student

import (
	"github.com/klinsc/model"
	"gorm.io/gorm"
)

type IRepository interface {
	Create(student *model.Student) error
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return &Repository{DB: db}
}

func (repo Repository) Create(student *model.Student) error {
	result := repo.DB.Create(student)

	return result.Error
}
