package loan

import "time"

type LoanRequest struct {
	UserId     int       `json:"user_id"`
	BookId     int       `json:"book_id"`
	DateReturn time.Time `json:"date_return"`
	Status     string    `json:"status"`
}
