package gamification

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{DB: db}
}

func (r *GormRepository) GetUserBadges(userID uuid.UUID) ([]UserBadge, error) {
	var userBadges []UserBadge
	err := r.DB.Where("user_id = ?", userID).Find(&userBadges).Error
	return userBadges, err
}

func (r *GormRepository) AwardBadge(userID, badgeID uuid.UUID) error {
	userBadge := UserBadge{
		UserID:   userID,
		BadgeID:  badgeID,
		EarnedAt: r.DB.NowFunc(),
	}
	return r.DB.Create(&userBadge).Error
}

func (r *GormRepository) GetUserStreaks(userID uuid.UUID) ([]Streak, error) {
	var streaks []Streak
	err := r.DB.Where("user_id = ?", userID).Find(&streaks).Error
	return streaks, err
}

func (r *GormRepository) UpdateStreak(userID uuid.UUID, streakType string, date string) error {
	// Implementation placeholder
	return nil
}

func (r *GormRepository) GetUserAchievements(userID uuid.UUID) ([]UserAchievement, error) {
	var userAchievements []UserAchievement
	err := r.DB.Where("user_id = ?", userID).Find(&userAchievements).Error
	return userAchievements, err
}

func (r *GormRepository) AwardAchievement(userID, achievementID uuid.UUID) error {
	userAchievement := UserAchievement{
		UserID:        userID,
		AchievementID: achievementID,
		EarnedAt:      r.DB.NowFunc(),
	}
	return r.DB.Create(&userAchievement).Error
}

func (r *GormRepository) GetLeaderboard(period string, limit int) ([]LeaderboardEntry, error) {
	var entries []LeaderboardEntry
	err := r.DB.Where("period = ?", period).Order("score DESC").Limit(limit).Find(&entries).Error
	return entries, err
}

func (r *GormRepository) GetUserLeaderboardEntry(userID uuid.UUID, period string) (*LeaderboardEntry, error) {
	var entry LeaderboardEntry
	err := r.DB.Where("user_id = ? AND period = ?", userID, period).First(&entry).Error
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *GormRepository) GetUserRewards(userID uuid.UUID) ([]UserReward, error) {
	var rewards []UserReward
	err := r.DB.Where("user_id = ?", userID).Find(&rewards).Error
	return rewards, err
}

func (r *GormRepository) ClaimReward(userID, rewardID uuid.UUID) error {
	userReward := UserReward{
		UserID:    userID,
		RewardID:  rewardID,
		ClaimedAt: r.DB.NowFunc(),
	}
	return r.DB.Create(&userReward).Error
}
