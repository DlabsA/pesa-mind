package user

import (
	"pesa-mind/internal/infrastructure/utils"
)

type User struct {
	utils.BaseModel
	Email        string   `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string   `gorm:"not null" json:"-"`
	Profile      *Profile `gorm:"foreignkey:UserID;references:ID" json:"profile,omitempty"`
}

type Profile struct {
	utils.BaseModel
	UserID   utils.UUID `gorm:"type:uuid;uniqueIndex;not null" json:"user_id"`
	User     *User      `gorm:"foreignkey:UserID;references:ID" json:"user,omitempty"`
	Username string     `gorm:"uniqueIndex;not null" json:"username"`
	Type     string     `gorm:"not null" json:"type"` // Enterprise, Premium, Free
	Balance  float64    `gorm:"not null;default:0" json:"balance"`
}
