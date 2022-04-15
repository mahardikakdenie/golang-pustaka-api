package entity

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
	UserId      int         `json:"author_id"`
	User        User        `json:"author"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `json:"deleted_at"`
}
