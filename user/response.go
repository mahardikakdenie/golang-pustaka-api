package user

import (
	"pustaka-api/entity"
	"time"
)

type UserResponse struct {
	Id        int                          `json:"id"`
	Name      string                       `json:"name"`
	Email     string                       `json:"email"`
	Book      []entity.Book                `json:"book"`
	Token     []entity.AuthenticationToken `json:"authentication_token"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
