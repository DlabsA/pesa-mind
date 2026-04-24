package category

import (
	"pesa-mind/internal/domain/user"
	"pesa-mind/internal/infrastructure/utils"

	"github.com/google/uuid"
)

type ChannelDetails struct {
	utils.BaseModel
	UserID      uuid.UUID         `gorm:"type:uuid;not null;index" json:"user_id"`
	User        *user.User        `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Name        string            `gorm:"not null" json:"name,omitempty"`
	ChannelType *user.ChannelType `gorm:"not null" json:"channel_type,omitempty"`
	ChannelDesc string            `gorm:"null" json:"channel_desc,omitempty"`
	Description string            `gorm:"not null" json:"description,omitempty"`
	Status      bool              `gorm:"not null" json:"active"`
}
