package gamification

import "github.com/google/uuid"

type Repository interface {
	GetUserBadges(userID uuid.UUID) ([]UserBadge, error)
	AwardBadge(userID, badgeID uuid.UUID) error
	GetUserStreaks(userID uuid.UUID) ([]Streak, error)
	UpdateStreak(userID uuid.UUID, streakType string, date string) error
	GetUserAchievements(userID uuid.UUID) ([]UserAchievement, error)
	AwardAchievement(userID, achievementID uuid.UUID) error
	GetLeaderboard(period string, limit int) ([]LeaderboardEntry, error)
	GetUserLeaderboardEntry(userID uuid.UUID, period string) (*LeaderboardEntry, error)
	GetUserRewards(userID uuid.UUID) ([]UserReward, error)
	ClaimReward(userID, rewardID uuid.UUID) error
}
