package gamification

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockRepo struct{}

func (m *mockRepo) GetUserBadges(userID uuid.UUID) ([]UserBadge, error) {
	return []UserBadge{{ID: uuid.New(), UserID: userID, BadgeID: uuid.New(), EarnedAt: time.Now()}}, nil
}
func (m *mockRepo) AwardBadge(userID, badgeID uuid.UUID) error { return nil }
func (m *mockRepo) GetUserStreaks(userID uuid.UUID) ([]Streak, error) {
	return []Streak{{ID: uuid.New(), UserID: userID, Type: "daily", Count: 5, LastDate: time.Now(), CreatedAt: time.Now(), UpdatedAt: time.Now()}}, nil
}
func (m *mockRepo) UpdateStreak(userID uuid.UUID, streakType string, date string) error { return nil }
func (m *mockRepo) GetUserAchievements(userID uuid.UUID) ([]UserAchievement, error) {
	return []UserAchievement{{ID: uuid.New(), UserID: userID, AchievementID: uuid.New(), EarnedAt: time.Now()}}, nil
}
func (m *mockRepo) AwardAchievement(userID, achievementID uuid.UUID) error { return nil }
func (m *mockRepo) GetLeaderboard(period string, limit int) ([]LeaderboardEntry, error) {
	return []LeaderboardEntry{{ID: uuid.New(), UserID: uuid.New(), Score: 100, Rank: 1, Period: period, CreatedAt: time.Now(), UpdatedAt: time.Now()}}, nil
}
func (m *mockRepo) GetUserLeaderboardEntry(userID uuid.UUID, period string) (*LeaderboardEntry, error) {
	return &LeaderboardEntry{ID: uuid.New(), UserID: userID, Score: 100, Rank: 1, Period: period, CreatedAt: time.Now(), UpdatedAt: time.Now()}, nil
}
func (m *mockRepo) GetUserRewards(userID uuid.UUID) ([]UserReward, error) {
	return []UserReward{{ID: uuid.New(), UserID: userID, RewardID: uuid.New(), ClaimedAt: time.Now()}}, nil
}
func (m *mockRepo) ClaimReward(userID, rewardID uuid.UUID) error { return nil }

func TestService_GetUserBadges(t *testing.T) {
	s := NewService(&mockRepo{})
	badges, err := s.GetUserBadges(uuid.New())
	assert.NoError(t, err)
	assert.NotEmpty(t, badges)
}

func TestService_GetUserStreaks(t *testing.T) {
	s := NewService(&mockRepo{})
	streaks, err := s.GetUserStreaks(uuid.New())
	assert.NoError(t, err)
	assert.NotEmpty(t, streaks)
}

func TestService_GetUserAchievements(t *testing.T) {
	s := NewService(&mockRepo{})
	a, err := s.GetUserAchievements(uuid.New())
	assert.NoError(t, err)
	assert.NotEmpty(t, a)
}

func TestService_GetLeaderboard(t *testing.T) {
	s := NewService(&mockRepo{})
	entries, err := s.GetLeaderboard("weekly", 10)
	assert.NoError(t, err)
	assert.NotEmpty(t, entries)
}

func TestService_GetUserRewards(t *testing.T) {
	s := NewService(&mockRepo{})
	r, err := s.GetUserRewards(uuid.New())
	assert.NoError(t, err)
	assert.NotEmpty(t, r)
}

func TestService_ClaimReward(t *testing.T) {
	s := NewService(&mockRepo{})
	err := s.ClaimReward(uuid.New(), uuid.New())
	assert.NoError(t, err)
}
