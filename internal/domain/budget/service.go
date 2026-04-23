package budget

import (
	"pesa-mind/internal/infrastructure/utils"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// Budget methods
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

func (s *Service) GetByID(id uuid.UUID) (*Budget, error) {
	return s.repo.FindByID(id)
}

func (s *Service) Update(b *Budget) error {
	return s.repo.Update(b)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

// MonthlyBudget methods
func (s *Service) CreateMonthlyBudget(userID, yearlyBudgetID uuid.UUID, month time.Month, year int64, transactions []BudgetTransaction) (*MonthlyBudget, error) {
	mb := &MonthlyBudget{
		UserID:             userID,
		YearlyBudgetID:     yearlyBudgetID,
		Month:              month,
		Year:               year,
		BudgetTransactions: transactions,
	}
	// Calculate totals from budget transactions
	s.calculateMonthlyBudgetTotals(mb)
	if err := s.repo.CreateMonthlyBudget(mb); err != nil {
		return nil, err
	}
	return mb, nil
}

func (s *Service) GetMonthlyBudgetByID(id uuid.UUID) (*MonthlyBudget, error) {
	return s.repo.FindMonthlyBudgetByID(id)
}

func (s *Service) GetMonthlyBudgetsByYearlyBudgetID(yearlyBudgetID uuid.UUID) ([]MonthlyBudget, error) {
	return s.repo.FindMonthlyBudgetsByYearlyBudgetID(yearlyBudgetID)
}

func (s *Service) GetMonthlyBudgetByUserIDAndMonthYear(userID uuid.UUID, month time.Month, year int64) (*MonthlyBudget, error) {
	return s.repo.FindMonthlyBudgetByUserIDAndMonthYear(userID, month, year)
}

func (s *Service) GetMonthlyBudgetsByUserID(userID uuid.UUID, limit, offset int) ([]MonthlyBudget, error) {
	return s.repo.FindMonthlyBudgetsByUserID(userID, limit, offset)
}

func (s *Service) UpdateMonthlyBudget(mb *MonthlyBudget) error {
	// Calculate totals from budget transactions
	s.calculateMonthlyBudgetTotals(mb)
	return s.repo.UpdateMonthlyBudget(mb)
}

func (s *Service) DeleteMonthlyBudget(id uuid.UUID) error {
	return s.repo.DeleteMonthlyBudget(id)
}

// YearlyBudget methods
func (s *Service) CreateYearlyBudget(userID uuid.UUID, year int64, transactions []BudgetTransaction) (*YearlyBudget, error) {
	yb := &YearlyBudget{
		UserID:             userID,
		Year:               year,
		BudgetTransactions: transactions,
	}
	// Calculate totals from budget transactions
	s.calculateYearlyBudgetTotals(yb)
	if err := s.repo.CreateYearlyBudget(yb); err != nil {
		return nil, err
	}
	return yb, nil
}

func (s *Service) GetYearlyBudgetByID(id uuid.UUID) (*YearlyBudget, error) {
	return s.repo.FindYearlyBudgetByID(id)
}

func (s *Service) GetYearlyBudgetsByUserID(userID uuid.UUID, limit, offset int) ([]YearlyBudget, error) {
	return s.repo.FindYearlyBudgetsByUserID(userID, limit, offset)
}

func (s *Service) UpdateYearlyBudget(yb *YearlyBudget) error {
	// Calculate totals from budget transactions
	s.calculateYearlyBudgetTotals(yb)
	return s.repo.UpdateYearlyBudget(yb)
}

func (s *Service) DeleteYearlyBudget(id uuid.UUID) error {
	return s.repo.DeleteYearlyBudget(id)
}

// Helper methods for calculating totals
func (s *Service) calculateMonthlyBudgetTotals(mb *MonthlyBudget) {
	var totalIncome uint64
	var totalExpenditures uint64
	var totalSavings uint64

	if mb.BudgetTransactions != nil {
		for _, bt := range mb.BudgetTransactions {
			switch bt.Type {
			case utils.TransactionTypeIncome:
				totalIncome += uint64(bt.Amount)
			case utils.TransactionTypeExpense:
				totalExpenditures += uint64(bt.Amount)
			case utils.TransactionTypeSaving:
				totalSavings += uint64(bt.Amount)
			}
		}
	}

	mb.TotalIncome = totalIncome
	mb.TotalExpenditures = totalExpenditures
	mb.TotalSavings = totalSavings
	mb.TotalTransactions = totalIncome - totalExpenditures - totalSavings
}

func (s *Service) calculateYearlyBudgetTotals(yb *YearlyBudget) {
	var totalIncome uint64
	var totalExpenditures uint64
	var totalSavings uint64

	if yb.BudgetTransactions != nil {
		for _, bt := range yb.BudgetTransactions {
			switch bt.Type {
			case utils.TransactionTypeIncome:
				totalIncome += uint64(bt.Amount)
			case utils.TransactionTypeExpense:
				totalExpenditures += uint64(bt.Amount)
			case utils.TransactionTypeSaving:
				totalSavings += uint64(bt.Amount)
			}
		}
	}

	yb.TotalIncome = totalIncome
	yb.TotalExpenditures = totalExpenditures
	yb.TotalSavings = totalSavings
	yb.TotalTransactions = totalIncome - totalExpenditures - totalSavings
}

func unixToTime(ts int64) time.Time {
	return time.Unix(ts, 0).UTC()
}
