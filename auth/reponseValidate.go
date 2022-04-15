package auth

import "time"

type ReponseAuth struct {
	ID          int       `json:"id"`
	UserId      int       `json:"user_id"`
	Email       string    `json:"email"`
	GeneratedAt time.Time `json:"generated_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}
