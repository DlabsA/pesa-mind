package account

import "github.com/google/uuid"

type Service struct {
	repo AccountRepository
}

func NewService(repo AccountRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(userID uuid.UUID, name, accType, currency string) (*Account, error) {
	account := &Account{
		ID:       uuid.New(),
		UserID:   userID,
		Name:     name,
		Type:     accType,
		Currency: currency,
	}
	if err := s.repo.Create(account); err != nil {
		return nil, err
	}
	return account, nil
}

func (s *Service) GetByID(id uuid.UUID) (*Account, error) {
	return s.repo.FindByID(id)
}

func (s *Service) GetByUserID(userID uuid.UUID) ([]*Account, error) {
	return s.repo.FindByUserID(userID)
}
