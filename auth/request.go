package auth

import "time"

type AuthRequest struct {
	ID          int       `json:"id"`
	UserId      int       `json:"user_id"`
	AuthToken   string    `json:"auth_token"`
	GeneratedAt time.Time `json:"generated_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	AuthType    string    `json:"auth_type"`
}
