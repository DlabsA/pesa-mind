package analytics

import (
	"pesa-mind/internal/domain/budget"
	"pesa-mind/internal/domain/savingsgoal"
	"pesa-mind/internal/domain/transaction"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repo Repository
}

// AnalyticsService aggregates repositories for analytics
type AnalyticsService struct {
	DB              *gorm.DB
	TransactionRepo transaction.TransactionRepository
	BudgetRepo      budget.Repository
	SavingsRepo     savingsgoal.Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func NewAnalyticsService(db *gorm.DB, txRepo transaction.TransactionRepository, bRepo budget.Repository, sRepo savingsgoal.Repository) *AnalyticsService {
	return &AnalyticsService{
		DB:              db,
		TransactionRepo: txRepo,
		BudgetRepo:      bRepo,
		SavingsRepo:     sRepo,
	}
}

func (s *Service) Create(userID uuid.UUID, typ, data, period string) (*AnalyticsSnapshot, error) {
	snap := &AnalyticsSnapshot{
		ID:     uuid.New(),
		UserID: userID,
		Type:   typ,
		Data:   data,
		Period: period,
	}
	if err := s.repo.Create(snap); err != nil {
		return nil, err
	}
	return snap, nil
}

func (s *Service) List(userID uuid.UUID, limit, offset int) ([]AnalyticsSnapshot, error) {
	return s.repo.FindByUserID(userID, limit, offset)
}

// TotalIncome returns the sum of all income transactions for a user in a date range
func (s *AnalyticsService) TotalIncome(userID uuid.UUID, from, to time.Time) (float64, error) {
	txs, err := s.TransactionRepo.FindByUserID(userID)
	if err != nil {
		return 0, err
	}
	total := 0.0
	for _, tx := range txs {
		if tx.Type == "income" && !tx.Date.Before(from) && !tx.Date.After(to) {
			total += tx.Amount
		}
	}
	return total, nil
}

// TotalExpenses returns the sum of all expense transactions for a user in a date range
func (s *AnalyticsService) TotalExpenses(userID uuid.UUID, from, to time.Time) (float64, error) {
	txs, err := s.TransactionRepo.FindByUserID(userID)
	if err != nil {
		return 0, err
	}
	total := 0.0
	for _, tx := range txs {
		if tx.Type == "expense" && !tx.Date.Before(from) && !tx.Date.After(to) {
			total += tx.Amount
		}
	}
	return total, nil
}

// BudgetUtilization returns the percentage of budget used for the current period
func (s *AnalyticsService) BudgetUtilization(userID uuid.UUID, now time.Time) (float64, error) {
	budgets, err := s.BudgetRepo.FindByUserID(userID, 100, 0)
	if err != nil {
		return 0, err
	}
	var totalBudget, totalSpent float64
	for _, b := range budgets {
		if now.After(b.StartDate) && now.Before(b.EndDate) {
			totalBudget += b.Amount
		}
	}
	txs, err := s.TransactionRepo.FindByUserID(userID)
	if err != nil {
		return 0, err
	}
	for _, tx := range txs {
		if tx.Type == "expense" && tx.Date.After(now.AddDate(0, -1, 0)) && tx.Date.Before(now) {
			totalSpent += tx.Amount
		}
	}
	if totalBudget == 0 {
		return 0, nil
	}
	return (totalSpent / totalBudget) * 100, nil
}

// SavingsProgress returns the percentage progress for all savings goals
func (s *AnalyticsService) SavingsProgress(userID uuid.UUID) (float64, error) {
	goals, err := s.SavingsRepo.FindByUserID(userID, 100, 0)
	if err != nil {
		return 0, err
	}
	var totalTarget, totalCurrent float64
	for _, g := range goals {
		totalTarget += g.Target
		totalCurrent += g.Current
	}
	if totalTarget == 0 {
		return 0, nil
	}
	return (totalCurrent / totalTarget) * 100, nil
}
