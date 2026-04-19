package user

import (
	"pesa-mind/internal/infrastructure/utils"
)

// User represents an application user. PasswordHash is omitted from JSON responses.
type User struct {
	utils.BaseModel
	Email        string `gorm:"uniqueIndex;not null" json:"email" validate:"required,email"`
	PasswordHash string `gorm:"not null" json:"-"`
	// One-to-one relationship to Profile. Cascade delete profile when user is deleted.
	Profile *Profile `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"profile,omitempty"`
}

// Profile contains user profile details and is linked to User via UserID.
type Profile struct {
	utils.BaseModel
	UserID   utils.UUID         `gorm:"type:uuid;uniqueIndex;not null" json:"user_id"`
	User     *User              `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Username string             `gorm:"uniqueIndex;not null" json:"username" validate:"required,min=3"`
	Type     string             `gorm:"not null;default:'Free'" json:"type"` // Enterprise, Premium, Free
	Balance  float64            `gorm:"not null;default:0" json:"balance"`
	Channels []FinancialChannel `gorm:"foreignKey:ProfileID;references:ID;constraint:OnDelete:CASCADE" json:"channels,omitempty"`
}

type FinancialChannel struct {
	utils.BaseModel
	ProfileID string      `gorm:"type:uuid;index" json:"profile_id"`
	Name      string      `gorm:"not null;default:''" json:"name"`
	Type      ChannelType `gorm:"column:channel_type;type:varchar(50);not null;default:'Cash'" json:"type"`
}

// ChannelDetails represents detailed information about a communication channel for a user
type ChannelDetails struct {
	utils.BaseModel
	UserID      utils.UUID  `gorm:"type:uuid;index;not null" json:"user_id"`
	ChannelType ChannelType `gorm:"type:varchar(50);not null" json:"channel_type"`
	Value       string      `gorm:"not null" json:"value"` // email, phone number, wallet address, etc.
	IsActive    bool        `gorm:"not null;default:true" json:"is_active"`
	IsPrimary   bool        `gorm:"not null;default:false" json:"is_primary"`
}
