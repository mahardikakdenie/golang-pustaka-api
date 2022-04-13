package user

import "pustaka-api/entity"

type Service interface {
	FindAll() ([]entity.User, error)
}

type service struct {
	repository Repository
}

func NewUserService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]entity.User, error) {
	return s.repository.FindAll()
}
