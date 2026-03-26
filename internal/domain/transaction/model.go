package transaction

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	AccountID   uuid.UUID  `gorm:"type:uuid;not null;index" json:"account_id"`
	CategoryID  uuid.UUID  `gorm:"type:uuid;not null;index" json:"category_id"`
	Amount      float64    `gorm:"not null" json:"amount"`
	Type        string     `gorm:"not null" json:"type"` // income/expense
	Note        string     `json:"note"`
	Description string     `json:"description,omitempty"`
	OccurredAt  *time.Time `json:"occurred_at,omitempty"`
	Date        time.Time  `gorm:"not null" json:"date"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
