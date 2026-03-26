package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct{ mock.Mock }

func (m *MockUserRepo) Create(user *User) error {
	args := m.Called(user)
	return args.Error(0)
}
func (m *MockUserRepo) FindByID(id uuid.UUID) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}
func (m *MockUserRepo) FindByEmail(email string) (*User, error) {
	args := m.Called(email)
	return args.Get(0).(*User), args.Error(1)
}
func (m *MockUserRepo) Update(user *User) error {
	args := m.Called(user)
	return args.Error(0)
}
func (m *MockUserRepo) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestRegister(t *testing.T) {
	repo := new(MockUserRepo)
	svc := NewService(repo)
	email := "test@example.com"
	passwordHash := "hashedpass"
	repo.On("Create", mock.AnythingOfType("*user.User")).Return(nil)
	user, err := svc.Register(email, passwordHash)
	assert.NoError(t, err)
	assert.Equal(t, email, user.Email)
}

func TestGetByID(t *testing.T) {
	repo := new(MockUserRepo)
	svc := NewService(repo)
	id := uuid.New()
	expected := &User{ID: id, Email: "id@example.com", PasswordHash: "hash"}
	repo.On("FindByID", id).Return(expected, nil)
	user, err := svc.GetByID(id)
	assert.NoError(t, err)
	assert.Equal(t, expected, user)
}

func TestGetByEmail(t *testing.T) {
	repo := new(MockUserRepo)
	svc := NewService(repo)
	email := "mail@example.com"
	expected := &User{ID: uuid.New(), Email: email, PasswordHash: "hash"}
	repo.On("FindByEmail", email).Return(expected, nil)
	user, err := svc.GetByEmail(email)
	assert.NoError(t, err)
	assert.Equal(t, expected, user)
}
