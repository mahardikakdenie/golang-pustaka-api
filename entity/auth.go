package entity

import (
	"time"

	"gorm.io/gorm"
)

type AuthenticationToken struct {
	gorm.Model
	UserId      int       `json:"user_id" gorm:"foreignkey:UserID"`
	User        User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AuthToken   string    `json:"auth_token"`
	AuthType    string    `json:"auth_type"`
	GeneratedAt time.Time `json:"generated_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
