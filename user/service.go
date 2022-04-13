package user

import "pustaka-api/entity"

type Service interface {
	FindAll() ([]entity.User, error)
	Create(userInput UserRequest) (entity.User, error)
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

func (s *service) Create(userInput UserRequest) (entity.User, error) {
	users, err := s.repository.Create(entity.User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: userInput.Password,
	})

	return users, err
}
