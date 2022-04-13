package book

import (
	"encoding/json"
	"pustaka-api/entity"
	"time"
)

type BookResponse struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Price       json.Number `json:"price"`
	Rating      json.Number `json:"rating"`
	Substitle   string      `json:"sub_title"`
	Description string      `json:"description"`
	UserId      int         `json:"author_id"`
	User        entity.User `json:"author"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
