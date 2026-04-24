package category

import (
	"pesa-mind/internal/domain/user"

	"github.com/google/uuid"
)

type Service struct {
	repo ChannelDetailsRepository
}

func NewService(repo ChannelDetailsRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(user *user.User, name, description, ChannelDesc string, channelType *user.ChannelType, status bool) (*ChannelDetails, error) {
	channelDetails := &ChannelDetails{
		UserID:      user.ID,
		User:        user,
		Name:        name,
		Description: description,
		ChannelDesc: ChannelDesc,
		ChannelType: channelType,
		Status:      status,
	}
	if err := s.repo.Create(channelDetails); err != nil {
		return nil, err
	}
	return channelDetails, nil
}

func (s *Service) GetByID(id uuid.UUID) (*ChannelDetails, error) {
	return s.repo.FindByID(id)
}

func (s *Service) GetByUserID(userID uuid.UUID) ([]*ChannelDetails, error) {
	return s.repo.FindByUserID(userID)
}

func (s *Service) GetByChannelType(channelType user.ChannelType) ([]*ChannelDetails, error) {
	return s.repo.FindByChannelType(channelType)
}

func (s *Service) GetByStatus(status bool) ([]*ChannelDetails, error) {
	return s.repo.FindByStatus(status)
}

func (s *Service) Update(channelDetails *ChannelDetails) error {
	return s.repo.Update(channelDetails)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
