package book

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Substitle   string      `json:"sub_title" binding:"required"`
	Rating      json.Number `json:"rating"`
	Description string      `json:"description"`
	UserId      int         `json:"user_id"`
	// User        user.User   `json:"user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
