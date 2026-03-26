package budget

import (
	"time"

	"github.com/google/uuid"
)

type Budget struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Name      string     `gorm:"not null" json:"name"`
	Amount    float64    `gorm:"not null" json:"amount"`
	Period    string     `gorm:"not null" json:"period"` // e.g. monthly, weekly
	StartDate time.Time  `gorm:"not null" json:"start_date"`
	EndDate   time.Time  `gorm:"not null" json:"end_date"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
