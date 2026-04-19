package transaction

import (
	"pesa-mind/internal/domain/category"
	"pesa-mind/internal/infrastructure/utils"
	"time"

	"pesa-mind/internal/domain/user"

	"github.com/google/uuid"
)

type Transaction struct {
	utils.BaseModel
	ProfileID        uuid.UUID                `gorm:"type:uuid;not null;index" json:"profile_id"`
	Profile          *user.Profile            `gorm:"foreignKey:ProfileID" json:"profile,omitempty"`
	CategoryID       uuid.UUID                `gorm:"type:uuid;not null;index" json:"category_id"`
	Amount           float64                  `gorm:"not null" json:"amount"`
	Type             string                   `gorm:"not null" json:"type"` // income/expense
	Note             string                   `json:"note"`
	ChannelDetailsID uuid.UUID                `gorm:"type:uuid;not null;index" json:"channel_details_id"`
	ChannelDetails   *category.ChannelDetails `gorm:"foreignKey:ChannelDetailsID;embedded;embeddedPrefix:channel_" json:"channel_details,omitempty"`
	Description      string                   `json:"description,omitempty"`
	OccurredAt       *time.Time               `json:"occurred_at,omitempty"`
	Date             time.Time                `gorm:"not null" json:"date"`
}
