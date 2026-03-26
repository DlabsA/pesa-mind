package category

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockCategoryRepo struct{ mock.Mock }

func (m *MockCategoryRepo) Create(category *Category) error {
	args := m.Called(category)
	return args.Error(0)
}
func (m *MockCategoryRepo) FindByID(id uuid.UUID) (*Category, error) {
	args := m.Called(id)
	return args.Get(0).(*Category), args.Error(1)
}
func (m *MockCategoryRepo) FindByUserID(userID uuid.UUID) ([]*Category, error) {
	args := m.Called(userID)
	return args.Get(0).([]*Category), args.Error(1)
}
func (m *MockCategoryRepo) Update(category *Category) error {
	args := m.Called(category)
	return args.Error(0)
}
func (m *MockCategoryRepo) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateCategory(t *testing.T) {
	repo := new(MockCategoryRepo)
	svc := NewService(repo)
	userID := uuid.New()
	name := "Test Category"
	typeStr := "expense"
	repo.On("Create", mock.AnythingOfType("*category.Category")).Return(nil)
	category, err := svc.Create(userID, name, typeStr, nil)
	assert.NoError(t, err)
	assert.Equal(t, name, category.Name)
	assert.Equal(t, typeStr, category.Type)
}
