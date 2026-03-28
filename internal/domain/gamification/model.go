package gamification

import (
	"time"

	"github.com/google/uuid"
)

type Badge struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string     `gorm:"not null;unique" json:"name"`
	Description string     `gorm:"not null" json:"description"`
	Icon        string     `json:"icon"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type UserBadge struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID   uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	BadgeID  uuid.UUID `gorm:"type:uuid;not null;index" json:"badge_id"`
	EarnedAt time.Time `gorm:"not null" json:"earned_at"`
}

type Streak struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Type      string    `gorm:"not null" json:"type"` // e.g. daily, weekly
	Count     int       `gorm:"not null" json:"count"`
	LastDate  time.Time `gorm:"not null" json:"last_date"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Achievement struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string     `gorm:"not null;unique" json:"name"`
	Description string     `gorm:"not null" json:"description"`
	Icon        string     `json:"icon"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type UserAchievement struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID        uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	AchievementID uuid.UUID `gorm:"type:uuid;not null;index" json:"achievement_id"`
	EarnedAt      time.Time `gorm:"not null" json:"earned_at"`
}

type LeaderboardEntry struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Score     int       `gorm:"not null" json:"score"`
	Rank      int       `gorm:"not null" json:"rank"`
	Period    string    `gorm:"not null" json:"period"` // e.g. weekly, monthly
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Reward struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string     `gorm:"not null;unique" json:"name"`
	Description string     `gorm:"not null" json:"description"`
	Points      int        `gorm:"not null" json:"points"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type UserReward struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	RewardID  uuid.UUID `gorm:"type:uuid;not null;index" json:"reward_id"`
	ClaimedAt time.Time `gorm:"not null" json:"claimed_at"`
}
