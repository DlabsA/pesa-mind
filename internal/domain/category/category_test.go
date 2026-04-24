package category

import (
	"pesa-mind/internal/domain/user"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockChannelDetailsRepo struct{ mock.Mock }

func (m *MockChannelDetailsRepo) Create(channelDetails *ChannelDetails) error {
	args := m.Called(channelDetails)
	return args.Error(0)
}

func (m *MockChannelDetailsRepo) FindByID(id uuid.UUID) (*ChannelDetails, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*ChannelDetails), args.Error(1)
}

func (m *MockChannelDetailsRepo) FindByUserID(userID uuid.UUID) ([]*ChannelDetails, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*ChannelDetails), args.Error(1)
}

func (m *MockChannelDetailsRepo) FindByChannelType(channelType user.ChannelType) ([]*ChannelDetails, error) {
	args := m.Called(channelType)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*ChannelDetails), args.Error(1)
}

func (m *MockChannelDetailsRepo) FindByStatus(status bool) ([]*ChannelDetails, error) {
	args := m.Called(status)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*ChannelDetails), args.Error(1)
}

func (m *MockChannelDetailsRepo) Update(channelDetails *ChannelDetails) error {
	args := m.Called(channelDetails)
	return args.Error(0)
}

func (m *MockChannelDetailsRepo) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

// Test Create ChannelDetails
func TestCreateChannelDetails(t *testing.T) {
	repo := new(MockChannelDetailsRepo)
	svc := NewService(repo)

	userID := uuid.New()
	testUser := &user.User{
		Email: "test@example.com",
	}
	testUser.ID = userID

	name := "Test Channel"
	description := "Test Description"
	channelType := user.ChannelTypeCash
	channelDesc := "Airtel Money"

	repo.On("Create", mock.AnythingOfType("*category.ChannelDetails")).Return(nil)

	result, err := svc.Create(testUser, name, description, channelDesc, &channelType, true)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, name, result.Name)
	assert.Equal(t, description, result.Description)
	assert.Equal(t, true, result.Status)
	repo.AssertExpectations(t)
}

// Test GetByID
func TestGetByIDChannelDetails(t *testing.T) {
	repo := new(MockChannelDetailsRepo)
	svc := NewService(repo)

	id := uuid.New()
	channelType := user.ChannelTypeCash
	testChannelDetails := &ChannelDetails{
		Name:        "Test Channel",
		Description: "Test Description",
		ChannelDesc: "Airtel Money",
		ChannelType: &channelType,
		Status:      true,
	}
	testChannelDetails.ID = id

	repo.On("FindByID", id).Return(testChannelDetails, nil)

	result, err := svc.GetByID(id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)
	assert.Equal(t, "Test Channel", result.Name)
	repo.AssertExpectations(t)
}

// Test GetByUserID
func TestGetByUserIDChannelDetails(t *testing.T) {
	repo := new(MockChannelDetailsRepo)
	svc := NewService(repo)

	userID := uuid.New()
	channelType := user.ChannelTypeCash
	testChannelDetails := []*ChannelDetails{
		{
			Name:        "Channel 1",
			Description: "Description 1",
			ChannelType: &channelType,
			ChannelDesc: "Airtel Money",
			Status:      true,
		},
		{
			Name:        "Channel 2",
			Description: "Description 2",
			ChannelType: &channelType,
			ChannelDesc: "Airtel Money",
			Status:      false,
		},
	}

	repo.On("FindByUserID", userID).Return(testChannelDetails, nil)

	result, err := svc.GetByUserID(userID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	assert.Equal(t, "Channel 1", result[0].Name)
	assert.Equal(t, "Channel 2", result[1].Name)
	repo.AssertExpectations(t)
}

// Test GetByChannelType
func TestGetByChannelTypeChannelDetails(t *testing.T) {
	repo := new(MockChannelDetailsRepo)
	svc := NewService(repo)

	channelType := user.ChannelTypeCash
	testChannelDetails := []*ChannelDetails{
		{
			Name:        "Cash Channel",
			Description: "Test Cash",
			ChannelType: &channelType,
			ChannelDesc: "Airtel Money",
			Status:      true,
		},
	}

	repo.On("FindByChannelType", channelType).Return(testChannelDetails, nil)

	result, err := svc.GetByChannelType(channelType)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Equal(t, "Cash Channel", result[0].Name)
	repo.AssertExpectations(t)
}

// Test GetByStatus
func TestGetByStatusChannelDetails(t *testing.T) {
	repo := new(MockChannelDetailsRepo)
	svc := NewService(repo)

	channelType := user.ChannelTypeCash
	testChannelDetails := []*ChannelDetails{
		{
			Name:        "Active Channel",
			Description: "Active",
			ChannelType: &channelType,
			ChannelDesc: "Airtel Money",
			Status:      true,
		},
	}

	repo.On("FindByStatus", true).Return(testChannelDetails, nil)

	result, err := svc.GetByStatus(true)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 1)
	assert.Equal(t, "Active Channel", result[0].Name)
	assert.True(t, result[0].Status)
	repo.AssertExpectations(t)
}

// Test Update
func TestUpdateChannelDetails(t *testing.T) {
	repo := new(MockChannelDetailsRepo)
	svc := NewService(repo)

	id := uuid.New()
	channelType := user.ChannelTypeCash
	channelDetails := &ChannelDetails{
		Name:        "Updated Channel",
		Description: "Updated Description",
		ChannelType: &channelType,
		ChannelDesc: "Airtel Money",
		Status:      false,
	}
	channelDetails.ID = id

	repo.On("Update", mock.MatchedBy(func(cd *ChannelDetails) bool {
		return cd.ID == id
	})).Return(nil)

	err := svc.Update(channelDetails)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

// Test Delete
func TestDeleteChannelDetails(t *testing.T) {
	repo := new(MockChannelDetailsRepo)
	svc := NewService(repo)

	id := uuid.New()
	repo.On("Delete", id).Return(nil)

	err := svc.Delete(id)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}
