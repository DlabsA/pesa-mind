package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTokenRepo struct{ mock.Mock }

func (m *MockTokenRepo) Create(token *RefreshToken) error {
	args := m.Called(token)
	return args.Error(0)
}
func (m *MockTokenRepo) FindByID(id uuid.UUID) (*RefreshToken, error) {
	args := m.Called(id)
	return args.Get(0).(*RefreshToken), args.Error(1)
}
func (m *MockTokenRepo) FindByUserID(userID uuid.UUID) ([]*RefreshToken, error) {
	args := m.Called(userID)
	return args.Get(0).([]*RefreshToken), args.Error(1)
}
func (m *MockTokenRepo) Revoke(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *MockTokenRepo) DeleteExpired(now time.Time) error {
	args := m.Called(now)
	return args.Error(0)
}

func TestCreateToken(t *testing.T) {
	repo := new(MockTokenRepo)
	svc := NewService(repo)
	token := &RefreshToken{ID: uuid.New(), UserID: uuid.New(), TokenHash: "hash", ExpiresAt: time.Now().Add(24 * time.Hour)}
	repo.On("Create", token).Return(nil)
	err := svc.tokenRepo.Create(token)
	assert.NoError(t, err)
}
