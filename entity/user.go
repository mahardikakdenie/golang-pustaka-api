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
	Book                []Book                `json:"book" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AuthenticationToken []AuthenticationToken `json:"authentication_token" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Loan                []Loan                `json:"loan" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
