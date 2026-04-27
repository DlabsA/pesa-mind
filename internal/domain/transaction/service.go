package transaction

import (
	"pesa-mind/internal/domain/category"
	"pesa-mind/internal/domain/gamification"
	"pesa-mind/internal/domain/user"

	"github.com/google/uuid"
)

type Service struct {
	repo         TransactionRepository
	Gamification *gamification.Service
	User         *user.Service
	Category     *category.Service
}

func NewService(repo TransactionRepository, gamification *gamification.Service, user *user.Service, category *category.Service) *Service {
	return &Service{repo: repo, Gamification: gamification, User: user, Category: category}
}

func (s *Service) Create(user *user.User, channelDetails *category.ChannelDetails, amount float64, txType, note string) (*Transaction, error) {
	tx := &Transaction{
		UserID:           user.ID,
		User:             user,
		ChannelDetailsID: channelDetails.ID,
		ChannelDetails:   channelDetails,
		Amount:           amount,
		Type:             txType,
		Note:             note,
	}
	if err := s.repo.Create(tx); err != nil {
		return nil, err
	}
	// Auto-award badge for first transaction
	//if s.Gamification != nil && user != nil {
	//	_ = s.Gamification.CheckAndAwardBadges(user.ID, "first_transaction")
	//}
	return tx, nil
}

func (s *Service) GetByID(id uuid.UUID) (*Transaction, error) {
	return s.repo.FindByID(id)
}

func (s *Service) GetByUserID(userID uuid.UUID) ([]*Transaction, error) {
	return s.repo.FindByUserID(userID)
}

func (s *Service) GetByUserIDAndType(userID uuid.UUID, txType string) ([]*Transaction, error) {
	return s.repo.FindByUserIDAndType(userID, txType)
}

// CreateTransactionFromAutomation allows automation to create a transaction for a user
func (s *Service) CreateTransactionFromAutomation(user *user.User, amount float64, channelDetails *category.ChannelDetails, description string) (*Transaction, error) {
	tx := &Transaction{
		UserID:           user.ID,
		User:             user,
		Amount:           amount,
		ChannelDetailsID: channelDetails.ID,
		ChannelDetails:   channelDetails,
		Description:      description,
	}
	if err := s.repo.Create(tx); err != nil {
		return nil, err
	}
	return tx, nil
}
