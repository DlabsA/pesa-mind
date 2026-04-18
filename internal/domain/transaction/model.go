package transaction

import (
	"pesa-mind/internal/infrastructure/utils"
	"time"

	"pesa-mind/internal/domain/user"

	"github.com/google/uuid"
)

type Transaction struct {
	utils.BaseModel
	Profile     *user.Profile    `gorm:"foreignKey:ProfileID" json:"profile,omitempty"`
	CategoryID  uuid.UUID        `gorm:"type:uuid;not null;index" json:"category_id"`
	Amount      float64          `gorm:"not null" json:"amount"`
	Type        string           `gorm:"not null" json:"type"` // income/expense
	Note        string           `json:"note"`
	Channel     user.ChannelType `gorm:"not null" json:"channel"`
	Description string           `json:"description,omitempty"`
	OccurredAt  *time.Time       `json:"occurred_at,omitempty"`
	Date        time.Time        `gorm:"not null" json:"date"`
}
