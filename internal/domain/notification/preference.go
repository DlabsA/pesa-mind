package notification

import (
	"github.com/google/uuid"
	"time"
)

type Preference struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	InApp  bool      `gorm:"default:true" json:"in_app"`
	Push   bool      `gorm:"default:true" json:"push"`
	Email  bool      `gorm:"default:false" json:"email"`
	// Future extensibility: add more channels as needed
	// SMS      bool      `gorm:"default:false" json:"sms"`
	// WhatsApp bool      `gorm:"default:false" json:"whatsapp"`
	// Custom   map[string]bool `gorm:"-" json:"custom,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
