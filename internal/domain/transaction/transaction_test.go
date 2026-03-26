package transaction

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionRepo struct{ mock.Mock }

func (m *MockTransactionRepo) Create(tx *Transaction) error {
	args := m.Called(tx)
	return args.Error(0)
}
func (m *MockTransactionRepo) FindByID(id uuid.UUID) (*Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(*Transaction), args.Error(1)
}
func (m *MockTransactionRepo) FindByUserID(userID uuid.UUID) ([]*Transaction, error) {
	args := m.Called(userID)
	return args.Get(0).([]*Transaction), args.Error(1)
}
func (m *MockTransactionRepo) Update(tx *Transaction) error {
	args := m.Called(tx)
	return args.Error(0)
}
func (m *MockTransactionRepo) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateTransaction(t *testing.T) {
	repo := new(MockTransactionRepo)
	svc := NewService(repo)
	userID := uuid.New()
	accountID := uuid.New()
	categoryID := uuid.New()
	amount := 100.0
	typeStr := "expense"
	note := "Test"
	date := time.Now().Unix()
	repo.On("Create", mock.AnythingOfType("*transaction.Transaction")).Return(nil)
	tx, err := svc.Create(userID, accountID, categoryID, amount, typeStr, note, date)
	assert.NoError(t, err)
	assert.Equal(t, amount, tx.Amount)
	assert.Equal(t, typeStr, tx.Type)
	assert.Equal(t, note, tx.Note)
}
