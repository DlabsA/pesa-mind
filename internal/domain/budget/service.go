package budget

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(userID uuid.UUID, name string, amount float64, period string, start, end int64) (*Budget, error) {
	b := &Budget{
		ID:        uuid.New(),
		UserID:    userID,
		Name:      name,
		Amount:    amount,
		Period:    period,
		StartDate: unixToTime(start),
		EndDate:   unixToTime(end),
	}
	if err := s.repo.Create(b); err != nil {
		return nil, err
	}
	return b, nil
}

func (s *Service) List(userID uuid.UUID, limit, offset int) ([]Budget, error) {
	return s.repo.FindByUserID(userID, limit, offset)
}

func unixToTime(ts int64) time.Time {
	return time.Unix(ts, 0).UTC()
}
