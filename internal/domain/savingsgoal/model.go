package savingsgoal

import (
	"time"

	"github.com/google/uuid"
)

type SavingsGoal struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Title     string     `gorm:"not null" json:"title"`
	Target    float64    `gorm:"not null" json:"target"`
	Current   float64    `gorm:"not null;default:0" json:"current"`
	Deadline  *time.Time `json:"deadline,omitempty"`
	AutoSave  bool       `gorm:"not null;default:false" json:"auto_save"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
