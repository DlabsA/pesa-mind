package savingsgoal

import (
	"github.com/google/uuid"
	"pesa-mind/internal/domain/gamification"
	"time"
)

type Service struct {
	repo         Repository
	Gamification *gamification.Service
}

func NewService(repo Repository, gamification *gamification.Service) *Service {
	return &Service{repo: repo, Gamification: gamification}
}

func (s *Service) Create(userID uuid.UUID, title string, target float64, deadline *int64, autoSave bool) (*SavingsGoal, error) {
	var deadlineTime *time.Time
	if deadline != nil {
		dt := time.Unix(*deadline, 0).UTC()
		deadlineTime = &dt
	}
	goal := &SavingsGoal{
		ID:       uuid.New(),
		UserID:   userID,
		Title:    title,
		Target:   target,
		Current:  0,
		Deadline: deadlineTime,
		AutoSave: autoSave,
	}
	if err := s.repo.Create(goal); err != nil {
		return nil, err
	}
	return goal, nil
}

func (s *Service) List(userID uuid.UUID, limit, offset int) ([]SavingsGoal, error) {
	return s.repo.FindByUserID(userID, limit, offset)
}

func (s *Service) CompleteGoal(userID, goalID uuid.UUID) error {
	// ...existing logic to mark goal as completed...
	if s.Gamification != nil {
		_ = s.Gamification.CheckAndAwardBadges(userID, "savings_goal_completed")
		_ = s.Gamification.CheckAndAwardAchievements(userID, "goal_achieved")
	}
	return nil
}
