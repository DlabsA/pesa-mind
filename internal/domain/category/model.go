package category

import (
	"pesa-mind/internal/domain/user"
	"pesa-mind/internal/infrastructure/utils"
)

type ChannelDetails struct {
	utils.BaseModel
	User        *user.User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Name        string            `gorm:"not null" json:"name,omitempty"`
	ChannelType *user.ChannelType `gorm:"not null" json:"channel_type,omitempty"`
	Description string            `gorm:"not null" json:"description,omitempty"`
	Status      bool              `gorm:"not null" json:"active"`
}
