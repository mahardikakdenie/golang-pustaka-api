package user

import (
	"pustaka-api/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Preload("Book.User").Find(&users).Error

	return users, err
}
