package gamification

import (
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUserBadges(userID uuid.UUID) ([]UserBadge, error) {
	return s.repo.GetUserBadges(userID)
}

func (s *Service) AwardBadge(userID, badgeID uuid.UUID) error {
	return s.repo.AwardBadge(userID, badgeID)
}

func (s *Service) GetUserStreaks(userID uuid.UUID) ([]Streak, error) {
	return s.repo.GetUserStreaks(userID)
}

func (s *Service) UpdateStreak(userID uuid.UUID, streakType string, date string) error {
	return s.repo.UpdateStreak(userID, streakType, date)
}

func (s *Service) GetUserAchievements(userID uuid.UUID) ([]UserAchievement, error) {
	return s.repo.GetUserAchievements(userID)
}

func (s *Service) AwardAchievement(userID, achievementID uuid.UUID) error {
	return s.repo.AwardAchievement(userID, achievementID)
}

func (s *Service) GetLeaderboard(period string, limit int) ([]LeaderboardEntry, error) {
	return s.repo.GetLeaderboard(period, limit)
}

func (s *Service) GetUserLeaderboardEntry(userID uuid.UUID, period string) (*LeaderboardEntry, error) {
	return s.repo.GetUserLeaderboardEntry(userID, period)
}

func (s *Service) GetUserRewards(userID uuid.UUID) ([]UserReward, error) {
	return s.repo.GetUserRewards(userID)
}

func (s *Service) ClaimReward(userID, rewardID uuid.UUID) error {
	return s.repo.ClaimReward(userID, rewardID)
}

func (s *Service) CheckAndAwardBadges(userID uuid.UUID, action string) error {
	badges, err := s.repo.GetUserBadges(userID)
	if err != nil {
		return err
	}
	badgeIDs := map[uuid.UUID]bool{}
	for _, b := range badges {
		badgeIDs[b.BadgeID] = true
	}
	// First Transaction Badge
	firstTxBadge := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	if action == "first_transaction" && !badgeIDs[firstTxBadge] {
		if err := s.repo.AwardBadge(userID, firstTxBadge); err != nil {
			return err
		}
	}
	// 7-Day Streak Badge
	streakBadge := uuid.MustParse("33333333-3333-3333-3333-333333333333")
	if action == "7_day_streak" && !badgeIDs[streakBadge] {
		if err := s.repo.AwardBadge(userID, streakBadge); err != nil {
			return err
		}
	}
	// Savings Goal Completed Badge
	goalBadge := uuid.MustParse("44444444-4444-4444-4444-444444444444")
	if action == "savings_goal_completed" && !badgeIDs[goalBadge] {
		if err := s.repo.AwardBadge(userID, goalBadge); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) CheckAndAwardAchievements(userID uuid.UUID, action string) error {
	achievements, err := s.repo.GetUserAchievements(userID)
	if err != nil {
		return err
	}
	achievementIDs := map[uuid.UUID]bool{}
	for _, a := range achievements {
		achievementIDs[a.AchievementID] = true
	}
	// Goal Achiever Achievement
	goalAchiever := uuid.MustParse("22222222-2222-2222-2222-222222222222")
	if action == "goal_achieved" && !achievementIDs[goalAchiever] {
		if err := s.repo.AwardAchievement(userID, goalAchiever); err != nil {
			return err
		}
	}
	// Super Saver Achievement
	superSaver := uuid.MustParse("55555555-5555-5555-5555-555555555555")
	if action == "super_saver" && !achievementIDs[superSaver] {
		if err := s.repo.AwardAchievement(userID, superSaver); err != nil {
			return err
		}
	}
	return nil
}
