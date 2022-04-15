package auth

import (
	"pustaka-api/entity"
)

type Service interface {
	GeneratedToken(tokenRequest AuthRequest) (entity.AuthenticationToken, error)
	ValidateToken(token string) (entity.AuthenticationToken, error)
	FindByEmail(email string) (entity.User, error)
	Register(user UserRequest) (entity.User, error)
	Destroy(token string) (entity.AuthenticationToken, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GeneratedToken(tokenRequest AuthRequest) (entity.AuthenticationToken, error) {
	// var user entity.User
	var token entity.AuthenticationToken
	token.ID = uint(tokenRequest.ID)
	token.AuthToken = tokenRequest.AuthToken
	token.ExpiresAt = tokenRequest.ExpiresAt
	token.GeneratedAt = tokenRequest.GeneratedAt
	token.UserId = tokenRequest.UserId
	token.AuthType = tokenRequest.AuthType

	_, err := s.repository.GeneratedToken(token)
	return token, err
}

func (s *service) FindByEmail(email string) (entity.User, error) {
	return s.repository.FindByEmail(email)
}

func (s *service) ValidateToken(token string) (entity.AuthenticationToken, error) {
	return s.repository.ValidateToken(token)
}

func (s *service) Register(request UserRequest) (entity.User, error) {
	var user entity.User
	user.ID = uint(request.ID)
	user.Name = request.Name
	user.Email = request.Email
	user.Password = request.Password

	return s.repository.Register(user)
}

func (s *service) Destroy(token string) (entity.AuthenticationToken, error) {
	tokens, _ := s.repository.ValidateToken(token)
	return s.repository.DestroyToken(tokens)
}
