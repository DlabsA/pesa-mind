package automation

import (
	"time"

	"github.com/google/uuid"
)

type SMSAutomation struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Pattern   string     `gorm:"not null" json:"pattern"` // Regex or keyword to match SMS
	Enabled   bool       `gorm:"default:true" json:"enabled"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
