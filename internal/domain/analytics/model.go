package analytics

import (
	"time"

	"github.com/google/uuid"
)

type AnalyticsSnapshot struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Type      string     `gorm:"not null" json:"type"`            // e.g. "trend", "report", "forecast"
	Data      string     `gorm:"type:jsonb;not null" json:"data"` // JSON-encoded analytics data
	Period    string     `gorm:"not null" json:"period"`          // e.g. "monthly", "weekly"
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
