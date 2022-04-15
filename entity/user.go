package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name                string                `json:"name"`
	Email               string                `json:"email" gorm:"index:,unique"`
	Password            string                `json:"password"`
	Book                []Book                `json:"book"`
	AuthenticationToken []AuthenticationToken `json:"authentication_token"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
