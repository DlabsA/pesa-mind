package account

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockAccountRepo struct{ mock.Mock }

func (m *MockAccountRepo) Create(account *Account) error {
	args := m.Called(account)
	return args.Error(0)
}
func (m *MockAccountRepo) FindByID(id uuid.UUID) (*Account, error) {
	args := m.Called(id)
	return args.Get(0).(*Account), args.Error(1)
}
func (m *MockAccountRepo) FindByUserID(userID uuid.UUID) ([]*Account, error) {
	args := m.Called(userID)
	return args.Get(0).([]*Account), args.Error(1)
}
func (m *MockAccountRepo) Update(account *Account) error {
	args := m.Called(account)
	return args.Error(0)
}
func (m *MockAccountRepo) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateAccount(t *testing.T) {
	repo := new(MockAccountRepo)
	svc := NewService(repo)
	userID := uuid.New()
	name := "Test Account"
	typeStr := "bank"
	currency := "USD"
	repo.On("Create", mock.AnythingOfType("*account.Account")).Return(nil)
	account, err := svc.Create(userID, name, typeStr, currency)
	assert.NoError(t, err)
	assert.Equal(t, name, account.Name)
	assert.Equal(t, typeStr, account.Type)
	assert.Equal(t, currency, account.Currency)
}
