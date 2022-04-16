package entity

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	gorm.Model
	Status     string    `json:"status"`
	DateReturn time.Time `json:"date_return"`
	UserId     int       `json:"user_id"`
	BookId     int       `json:"book_id"`
	User       User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Book       Book      `json:"book" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
