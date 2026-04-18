package transaction

import (
	"pesa-mind/internal/domain/user"
	"time"

	"pesa-mind/internal/domain/gamification"

	"github.com/google/uuid"
)

type Service struct {
	repo         TransactionRepository
	Gamification *gamification.Service
	Profile      *user.Service
}

func NewService(repo TransactionRepository, gamification *gamification.Service, profile *user.Service) *Service {
	return &Service{repo: repo, Gamification: gamification, Profile: profile}
}

func (s *Service) Create(profile *user.Profile, categoryID uuid.UUID, amount float64, txType, note string, date int64) (*Transaction, error) {
	tx := &Transaction{
		Profile:    profile,
		CategoryID: categoryID,
		Amount:     amount,
		Type:       txType,
		Note:       note,
		Date:       time.Unix(date, 0),
	}
	if err := s.repo.Create(tx); err != nil {
		return nil, err
	}
	// Auto-award badge for first transaction
	if s.Gamification != nil && profile != nil {
		_ = s.Gamification.CheckAndAwardBadges(profile.UserID, "first_transaction")
	}
	return tx, nil
}

func (s *Service) GetByID(id uuid.UUID) (*Transaction, error) {
	return s.repo.FindByID(id)
}

func (s *Service) GetByUserID(userID uuid.UUID) ([]*Transaction, error) {
	return s.repo.FindByUserID(userID)
}

// CreateTransactionFromAutomation allows automation to create a transaction for a user
func (s *Service) CreateTransactionFromAutomation(profile *user.Profile, amount float64, categoryID uuid.UUID, description string, occurredAt time.Time) (*Transaction, error) {
	tx := &Transaction{
		Profile:     profile,
		Amount:      amount,
		CategoryID:  categoryID,
		Description: description,
		OccurredAt:  &occurredAt,
	}
	if err := s.repo.Create(tx); err != nil {
		return nil, err
	}
	return tx, nil
}
