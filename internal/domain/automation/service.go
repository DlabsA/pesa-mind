package automation

import (
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateRule(userID uuid.UUID, pattern string) (*SMSAutomation, error) {
	rule := &SMSAutomation{
		UserID:  userID,
		Pattern: pattern,
		Enabled: true,
	}
	err := s.repo.Create(rule)
	return rule, err
}

func (s *Service) ListRules(userID uuid.UUID) ([]SMSAutomation, error) {
	return s.repo.FindByUserID(userID)
}

func (s *Service) UpdateRule(rule *SMSAutomation) error {
	return s.repo.Update(rule)
}

func (s *Service) DeleteRule(id uuid.UUID) error {
	return s.repo.Delete(id)
}
