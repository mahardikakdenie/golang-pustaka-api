package user

import "pustaka-api/entity"

type Service interface {
	FindAll() ([]entity.User, error)
	FindById(id int) (entity.User, error)
	Create(userInput UserRequest) (entity.User, error)
	Update(userInput UserRequest, id int) (entity.User, error)
	Destroy(id int) (entity.User, error)
	Login(email string, password string) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
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

func (s *service) FindById(id int) (entity.User, error) {
	return s.repository.FindById(id)
}

func (s *service) Update(userInput UserRequest, id int) (entity.User, error) {
	userId, err := s.repository.FindById(id)
	if err != nil {
	}
	userId.Name = userInput.Name
	userId.Email = userInput.Email
	userId.Password = userInput.Password

	userUpdate, err := s.repository.Update(userId)

	return userUpdate, err
}

func (s *service) Destroy(id int) (entity.User, error) {
	userById, _ := s.repository.FindById(id)
	user, err := s.repository.Destroy(userById)

	return user, err
}

func (s *service) Login(email string, password string) (entity.User, error) {
	// var users entity.User
	users, err := s.repository.Login(email, password)
	return users, err
}

func (s *service) FindByEmail(email string) (entity.User, error) {
	return s.repository.FindByEmail(email)
}
