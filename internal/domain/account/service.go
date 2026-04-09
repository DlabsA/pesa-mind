package account

import (
	"github.com/google/uuid"
)

type Service struct {
	repo AccountRepository
}

func NewService(repo AccountRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(userID uuid.UUID, name, accType, currency string) (*Account, error) {
	acc := &Account{
		UserID:   userID,
		Name:     name,
		Type:     accType,
		Currency: currency,
		Balance:  0.0,
	}
	if err := s.repo.Create(acc); err != nil {
		return nil, err
	}
	return acc, nil
}

func (s *Service) GetByID(id uuid.UUID) (*Account, error) {
	return s.repo.FindByID(id)
}

func (s *Service) GetByUserID(userID uuid.UUID) ([]*Account, error) {
	return s.repo.FindByUserID(userID)
}
