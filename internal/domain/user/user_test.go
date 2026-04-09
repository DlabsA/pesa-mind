package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct{ mock.Mock }

func (m *MockUserRepo) Create(userProfile UserProfile) error {
	args := m.Called(userProfile)
	return args.Error(0)
}
func (m *MockUserRepo) FindByID(id uuid.UUID) (*User, *Profile, error) {
	args := m.Called(id)
	user := args.Get(0)
	if user != nil {
		return user.(*User), nil, args.Error(1)
	}
	return nil, nil, args.Error(1)
}
func (m *MockUserRepo) FindByEmail(email string) (*User, *Profile, error) {
	args := m.Called(email)
	user := args.Get(0)
	if user != nil {
		return user.(*User), nil, args.Error(1)
	}
	return nil, nil, args.Error(1)
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
	repo.On("Create", mock.MatchedBy(func(up UserProfile) bool {
		return up.user.Email == email && up.username == email
	})).Return(nil)
	user, err := svc.Register(email, passwordHash)
	assert.NoError(t, err)
	assert.Equal(t, email, user.Email)
}

func TestGetByID(t *testing.T) {
	repo := new(MockUserRepo)
	svc := NewService(repo)
	id := uuid.New()
	expected := &User{Email: "id@example.com", PasswordHash: "hash"}
	expected.ID = id
	repo.On("FindByID", id).Return(expected, nil)
	user, _, err := svc.GetByID(id)
	assert.NoError(t, err)
	assert.Equal(t, expected, user)
}

func TestGetByEmail(t *testing.T) {
	repo := new(MockUserRepo)
	svc := NewService(repo)
	email := "mail@example.com"
	expected := &User{Email: email, PasswordHash: "hash"}
	expected.ID = uuid.New()
	repo.On("FindByEmail", email).Return(expected, nil)
	user, _, err := svc.GetByEmail(email)
	assert.NoError(t, err)
	assert.Equal(t, expected, user)
}
