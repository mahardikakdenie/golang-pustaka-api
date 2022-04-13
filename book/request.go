package book

import "encoding/json"

type BookRequest struct {
	Id          int         `json:"id"`
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required"`
	Substitle   string      `json:"sub_title" binding:"required"`
	Rating      json.Number `json:"rating" binding:"required"`
	Description string      `json:"description" binding:"required"`
	UserId      int         `json:"user_id"`
}
