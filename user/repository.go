package user

import (
	"pustaka-api/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.User, error)
	FindById(id int) (entity.User, error)
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Destroy(users entity.User) (entity.User, error)
	Login(email string, password string) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Preload("Book.User").Preload("AuthenticationToken").Find(&users).Error

	return users, err
}

func (r *repository) Create(users entity.User) (entity.User, error) {
	err := r.db.Create(&users).Error
	return users, err
}

func (r *repository) FindById(Id int) (entity.User, error) {
	var user entity.User
	err := r.db.Preload("Book.User").Preload("AuthenticationToken").Find(&user, Id).Error

	return user, err
}

func (r *repository) Update(users entity.User) (entity.User, error) {
	err := r.db.Save(&users).Error
	return users, err
}

func (r *repository) Destroy(users entity.User) (entity.User, error) {
	err := r.db.Delete(&users).Error

	return users, err
}

func (r *repository) Login(email string, password string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ? AND password = ?", email, password).Find(&user).Error
	return user, err
}

func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	return user, err
}
