package budget

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockRepo struct {
	budgets        []Budget
	monthlyBudgets []MonthlyBudget
	yearlyBudgets  []YearlyBudget
}

// Budget methods
func (m *mockRepo) Create(b *Budget) error {
	m.budgets = append(m.budgets, *b)
	return nil
}

func (m *mockRepo) FindByID(id uuid.UUID) (*Budget, error) {
	for _, b := range m.budgets {
		if b.ID == id {
			return &b, nil
		}
	}
	return nil, nil
}

func (m *mockRepo) FindByUserID(userID uuid.UUID, limit, offset int) ([]Budget, error) {
	var res []Budget
	for _, b := range m.budgets {
		if b.UserID == userID {
			res = append(res, b)
		}
	}
	return res, nil
}

func (m *mockRepo) Update(b *Budget) error    { return nil }
func (m *mockRepo) Delete(id uuid.UUID) error { return nil }

// MonthlyBudget methods
func (m *mockRepo) CreateMonthlyBudget(mb *MonthlyBudget) error {
	m.monthlyBudgets = append(m.monthlyBudgets, *mb)
	return nil
}

func (m *mockRepo) FindMonthlyBudgetByID(id uuid.UUID) (*MonthlyBudget, error) {
	for _, mb := range m.monthlyBudgets {
		if mb.ID == id {
			return &mb, nil
		}
	}
	return nil, nil
}

func (m *mockRepo) FindMonthlyBudgetsByYearlyBudgetID(yearlyBudgetID uuid.UUID) ([]MonthlyBudget, error) {
	var res []MonthlyBudget
	for _, mb := range m.monthlyBudgets {
		if mb.YearlyBudgetID == yearlyBudgetID {
			res = append(res, mb)
		}
	}
	return res, nil
}

func (m *mockRepo) FindMonthlyBudgetByUserIDAndMonthYear(userID uuid.UUID, month time.Month, year int64) (*MonthlyBudget, error) {
	for _, mb := range m.monthlyBudgets {
		if mb.UserID == userID && mb.Month == month && mb.Year == year {
			return &mb, nil
		}
	}
	return nil, nil
}

func (m *mockRepo) FindMonthlyBudgetsByUserID(userID uuid.UUID, limit, offset int) ([]MonthlyBudget, error) {
	var res []MonthlyBudget
	for _, mb := range m.monthlyBudgets {
		if mb.UserID == userID {
			res = append(res, mb)
		}
	}
	return res, nil
}

func (m *mockRepo) UpdateMonthlyBudget(mb *MonthlyBudget) error {
	for i, b := range m.monthlyBudgets {
		if b.ID == mb.ID {
			m.monthlyBudgets[i] = *mb
			return nil
		}
	}
	return nil
}

func (m *mockRepo) DeleteMonthlyBudget(id uuid.UUID) error {
	for i, mb := range m.monthlyBudgets {
		if mb.ID == id {
			m.monthlyBudgets = append(m.monthlyBudgets[:i], m.monthlyBudgets[i+1:]...)
			return nil
		}
	}
	return nil
}

// YearlyBudget methods
func (m *mockRepo) CreateYearlyBudget(yb *YearlyBudget) error {
	m.yearlyBudgets = append(m.yearlyBudgets, *yb)
	return nil
}

func (m *mockRepo) FindYearlyBudgetByID(id uuid.UUID) (*YearlyBudget, error) {
	for _, yb := range m.yearlyBudgets {
		if yb.ID == id {
			return &yb, nil
		}
	}
	return nil, nil
}

func (m *mockRepo) FindYearlyBudgetsByUserID(userID uuid.UUID, limit, offset int) ([]YearlyBudget, error) {
	var res []YearlyBudget
	for _, yb := range m.yearlyBudgets {
		if yb.UserID == userID {
			res = append(res, yb)
		}
	}
	return res, nil
}

func (m *mockRepo) UpdateYearlyBudget(yb *YearlyBudget) error {
	for i, b := range m.yearlyBudgets {
		if b.ID == yb.ID {
			m.yearlyBudgets[i] = *yb
			return nil
		}
	}
	return nil
}

func (m *mockRepo) DeleteYearlyBudget(id uuid.UUID) error {
	for i, yb := range m.yearlyBudgets {
		if yb.ID == id {
			m.yearlyBudgets = append(m.yearlyBudgets[:i], m.yearlyBudgets[i+1:]...)
			return nil
		}
	}
	return nil
}

// Budget Tests
func TestService_CreateAndList(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()
	b, err := svc.Create(userID, "Test Budget", 1000, "monthly", 1700000000, 1702592000)
	assert.NoError(t, err)
	assert.Equal(t, "Test Budget", b.Name)
	budgets, err := svc.List(userID, 10, 0)
	assert.NoError(t, err)
	assert.Len(t, budgets, 1)
}

func TestService_GetBudgetByID(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()

	b, err := svc.Create(userID, "Test Budget", 1000, "monthly", 1700000000, 1702592000)
	assert.NoError(t, err)

	retrieved, err := svc.GetByID(b.ID)
	assert.NoError(t, err)
	assert.Equal(t, b.ID, retrieved.ID)
	assert.Equal(t, "Test Budget", retrieved.Name)
}

func TestService_UpdateBudget(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()

	b, err := svc.Create(userID, "Test Budget", 1000, "monthly", 1700000000, 1702592000)
	assert.NoError(t, err)

	b.Name = "Updated Budget"
	err = svc.Update(b)
	assert.NoError(t, err)
}

func TestService_DeleteBudget(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()

	b, err := svc.Create(userID, "Test Budget", 1000, "monthly", 1700000000, 1702592000)
	assert.NoError(t, err)

	err = svc.Delete(b.ID)
	assert.NoError(t, err)
}

// MonthlyBudget Tests
func TestService_CreateMonthlyBudget_WithTotalsCalculation(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()
	yearlyBudgetID := uuid.New()

	transactions := []BudgetTransaction{
		{
			Name:   "Salary",
			Amount: 5000,
			Type:   "income",
		},
		{
			Name:   "Expenses",
			Amount: 2000,
			Type:   "expense",
		},
		{
			Name:   "Savings",
			Amount: 1000,
			Type:   "saving",
		},
	}

	mb, err := svc.CreateMonthlyBudget(userID, yearlyBudgetID, time.January, 2024, transactions)
	assert.NoError(t, err)
	assert.NotNil(t, mb)
	assert.Equal(t, uint64(5000), mb.TotalIncome)
	assert.Equal(t, uint64(2000), mb.TotalExpenditures)
	assert.Equal(t, uint64(1000), mb.TotalSavings)
	assert.Equal(t, uint64(2000), mb.TotalTransactions) // 5000 - 2000 - 1000
}

func TestService_GetMonthlyBudgetByID(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()
	yearlyBudgetID := uuid.New()

	transactions := []BudgetTransaction{}
	mb, err := svc.CreateMonthlyBudget(userID, yearlyBudgetID, time.February, 2024, transactions)
	assert.NoError(t, err)

	retrieved, err := svc.GetMonthlyBudgetByID(mb.ID)
	assert.NoError(t, err)
	assert.Equal(t, mb.ID, retrieved.ID)
	assert.Equal(t, time.February, retrieved.Month)
}

func TestService_GetMonthlyBudgetsByYearlyBudgetID(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()
	yearlyBudgetID := uuid.New()

	transactions := []BudgetTransaction{}

	mb1, err := svc.CreateMonthlyBudget(userID, yearlyBudgetID, time.January, 2024, transactions)
	assert.NoError(t, err)

	mb2, err := svc.CreateMonthlyBudget(userID, yearlyBudgetID, time.February, 2024, transactions)
	assert.NoError(t, err)

	retrieved, err := svc.GetMonthlyBudgetsByYearlyBudgetID(yearlyBudgetID)
	assert.NoError(t, err)
	assert.Len(t, retrieved, 2)
	assert.True(t, (retrieved[0].ID == mb1.ID || retrieved[0].ID == mb2.ID))
}

func TestService_GetMonthlyBudgetByUserIDAndMonthYear(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()
	yearlyBudgetID := uuid.New()

	transactions := []BudgetTransaction{}
	mb, err := svc.CreateMonthlyBudget(userID, yearlyBudgetID, time.March, 2024, transactions)
	assert.NoError(t, err)

	retrieved, err := svc.GetMonthlyBudgetByUserIDAndMonthYear(userID, time.March, 2024)
	assert.NoError(t, err)
	assert.Equal(t, mb.ID, retrieved.ID)
	assert.Equal(t, time.March, retrieved.Month)
	assert.Equal(t, int64(2024), retrieved.Year)
}

func TestService_GetMonthlyBudgetsByUserID(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()
	yearlyBudgetID := uuid.New()

	transactions := []BudgetTransaction{}

	mb1, err := svc.CreateMonthlyBudget(userID, yearlyBudgetID, time.April, 2024, transactions)
	assert.NoError(t, err)

	mb2, err := svc.CreateMonthlyBudget(userID, yearlyBudgetID, time.May, 2024, transactions)
	assert.NoError(t, err)

	retrieved, err := svc.GetMonthlyBudgetsByUserID(userID, 10, 0)
	assert.NoError(t, err)
	assert.Len(t, retrieved, 2)
	assert.True(t, (retrieved[0].ID == mb1.ID || retrieved[0].ID == mb2.ID))
}

func TestService_UpdateMonthlyBudget_WithTotalsRecalculation(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()
	yearlyBudgetID := uuid.New()

	oldTransactions := []BudgetTransaction{
		{Name: "Salary", Amount: 5000, Type: "income"},
	}

	mb, err := svc.CreateMonthlyBudget(userID, yearlyBudgetID, time.June, 2024, oldTransactions)
	assert.NoError(t, err)
	assert.Equal(t, uint64(5000), mb.TotalIncome)

	// Update with new transactions
	mb.BudgetTransactions = []BudgetTransaction{
		{Name: "Salary", Amount: 6000, Type: "income"},
		{Name: "Bonus", Amount: 1000, Type: "income"},
		{Name: "Expenses", Amount: 3000, Type: "expense"},
	}

	err = svc.UpdateMonthlyBudget(mb)
	assert.NoError(t, err)
	assert.Equal(t, uint64(7000), mb.TotalIncome)
	assert.Equal(t, uint64(3000), mb.TotalExpenditures)
}

func TestService_DeleteMonthlyBudget(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()
	yearlyBudgetID := uuid.New()

	transactions := []BudgetTransaction{}
	mb, err := svc.CreateMonthlyBudget(userID, yearlyBudgetID, time.July, 2024, transactions)
	assert.NoError(t, err)

	err = svc.DeleteMonthlyBudget(mb.ID)
	assert.NoError(t, err)

	retrieved, err := svc.GetMonthlyBudgetByID(mb.ID)
	assert.NoError(t, err)
	assert.Nil(t, retrieved)
}

// YearlyBudget Tests
func TestService_CreateYearlyBudget_WithTotalsCalculation(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()

	transactions := []BudgetTransaction{
		{Name: "Annual Salary", Amount: 60000, Type: "income"},
		{Name: "Annual Expenses", Amount: 24000, Type: "expense"},
		{Name: "Annual Savings", Amount: 12000, Type: "saving"},
	}

	yb, err := svc.CreateYearlyBudget(userID, 2024, transactions)
	assert.NoError(t, err)
	assert.NotNil(t, yb)
	assert.Equal(t, uint64(60000), yb.TotalIncome)
	assert.Equal(t, uint64(24000), yb.TotalExpenditures)
	assert.Equal(t, uint64(12000), yb.TotalSavings)
	assert.Equal(t, uint64(24000), yb.TotalTransactions) // 60000 - 24000 - 12000
}

func TestService_GetYearlyBudgetByID(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()

	transactions := []BudgetTransaction{}
	yb, err := svc.CreateYearlyBudget(userID, 2024, transactions)
	assert.NoError(t, err)

	retrieved, err := svc.GetYearlyBudgetByID(yb.ID)
	assert.NoError(t, err)
	assert.Equal(t, yb.ID, retrieved.ID)
	assert.Equal(t, int64(2024), retrieved.Year)
}

func TestService_GetYearlyBudgetsByUserID(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()

	transactions := []BudgetTransaction{}

	yb1, err := svc.CreateYearlyBudget(userID, 2023, transactions)
	assert.NoError(t, err)

	yb2, err := svc.CreateYearlyBudget(userID, 2024, transactions)
	assert.NoError(t, err)

	retrieved, err := svc.GetYearlyBudgetsByUserID(userID, 10, 0)
	assert.NoError(t, err)
	assert.Len(t, retrieved, 2)
	assert.True(t, (retrieved[0].ID == yb1.ID || retrieved[0].ID == yb2.ID))
}

func TestService_UpdateYearlyBudget_WithTotalsRecalculation(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()

	oldTransactions := []BudgetTransaction{
		{Name: "Annual Salary", Amount: 50000, Type: "income"},
	}

	yb, err := svc.CreateYearlyBudget(userID, 2024, oldTransactions)
	assert.NoError(t, err)
	assert.Equal(t, uint64(50000), yb.TotalIncome)

	// Update with new transactions
	yb.BudgetTransactions = []BudgetTransaction{
		{Name: "Annual Salary", Amount: 60000, Type: "income"},
		{Name: "Bonus", Amount: 10000, Type: "income"},
		{Name: "Annual Expenses", Amount: 30000, Type: "expense"},
		{Name: "Savings", Amount: 15000, Type: "saving"},
	}

	err = svc.UpdateYearlyBudget(yb)
	assert.NoError(t, err)
	assert.Equal(t, uint64(70000), yb.TotalIncome)
	assert.Equal(t, uint64(30000), yb.TotalExpenditures)
	assert.Equal(t, uint64(15000), yb.TotalSavings)
	assert.Equal(t, uint64(25000), yb.TotalTransactions) // 70000 - 30000 - 15000
}

func TestService_DeleteYearlyBudget(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()

	transactions := []BudgetTransaction{}
	yb, err := svc.CreateYearlyBudget(userID, 2024, transactions)
	assert.NoError(t, err)

	err = svc.DeleteYearlyBudget(yb.ID)
	assert.NoError(t, err)

	retrieved, err := svc.GetYearlyBudgetByID(yb.ID)
	assert.NoError(t, err)
	assert.Nil(t, retrieved)
}
