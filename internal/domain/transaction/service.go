package transaction

import (
	"time"

	"github.com/google/uuid"
	"pesa-mind/internal/domain/gamification"
)

type Service struct {
	repo         TransactionRepository
	Gamification *gamification.Service
}

func NewService(repo TransactionRepository, gamification *gamification.Service) *Service {
	return &Service{repo: repo, Gamification: gamification}
}

func (s *Service) Create(userID, accountID, categoryID uuid.UUID, amount float64, txType, note string, date int64) (*Transaction, error) {
	tx := &Transaction{
		ID:         uuid.New(),
		UserID:     userID,
		AccountID:  accountID,
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
	if s.Gamification != nil {
		_ = s.Gamification.CheckAndAwardBadges(userID, "first_transaction")
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
func (s *Service) CreateTransactionFromAutomation(userID uuid.UUID, amount float64, categoryID uuid.UUID, description string, occurredAt time.Time) (*Transaction, error) {
	tx := &Transaction{
		UserID:      userID,
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
