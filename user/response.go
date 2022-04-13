package user

import (
	"pustaka-api/entity"
	"time"
)

type UserResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Password  string              `json:"password" binding:"hidden"`
	Book      []entity.Book `json:"book"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
