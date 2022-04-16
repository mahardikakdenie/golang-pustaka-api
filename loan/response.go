package loan

import (
	"pustaka-api/entity"
	"time"
)

type LoanResponse struct {
	ID         uint        `json:"id"`
	Status     string      `json:"status"`
	DateReturn time.Time   `json:"date_return"`
	UserId     int         `json:"user_id"`
	BookId     int         `json:"book_id"`
	User       entity.User `json:"borrower" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Book       entity.Book `json:"book" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
