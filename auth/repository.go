package auth

import (
	"pustaka-api/entity"

	"gorm.io/gorm"
)

type Repository interface {
	GeneratedToken(token entity.AuthenticationToken) (entity.AuthenticationToken, error)
	ValidateToken(token string) (entity.AuthenticationToken, error)
	FindByToken(token string) (entity.AuthenticationToken, error)
	FindByEmail(email string) (entity.User, error)
	Register(user entity.User) (entity.User, error)
	DestroyToken(token entity.AuthenticationToken) (entity.AuthenticationToken, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GeneratedToken(token entity.AuthenticationToken) (entity.AuthenticationToken, error) {
	err := r.db.Create(&token).Error
	return token, err
}

func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	return user, err
}

func (r *repository) ValidateToken(token string) (entity.AuthenticationToken, error) {
	var tokenEntity entity.AuthenticationToken
	err := r.db.Preload("User").Where("auth_token = ?", token).Find(&tokenEntity).Error
	return tokenEntity, err
}

func (r *repository) FindByToken(token string) (entity.AuthenticationToken, error) {
	var tokenEntity entity.AuthenticationToken
	err := r.db.Where("auth_token = ?", token).Find(&tokenEntity).Error
	return tokenEntity, err
}

func (r *repository) Register(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) DestroyToken(token entity.AuthenticationToken) (entity.AuthenticationToken, error) {
	err := r.db.Delete(&token).Error
	return token, err
}
