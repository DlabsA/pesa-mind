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
	UserID           uuid.UUID                `gorm:"type:uuid;not null;index" json:"user_id"`
	User             *user.User               `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Amount           float64                  `gorm:"not null" json:"amount"`
	Type             string                   `gorm:"not null" json:"type"` // income/expense
	Note             string                   `json:"note"`
	ChannelDetailsID uuid.UUID                `gorm:"type:uuid;not null;index" json:"channel_details_id"`
	ChannelDetails   *category.ChannelDetails `gorm:"foreignKey:ChannelDetailsID" json:"channel_details,omitempty"`
	Description      string                   `json:"description,omitempty"`
	OccurredAt       *time.Time               `json:"occurred_at,omitempty"`
}
