package category

import (
	"pesa-mind/internal/domain/user"

	"github.com/google/uuid"
)

type ChannelDetailsRepository interface {
	Create(channelDetails *ChannelDetails) error
	FindByID(id uuid.UUID) (*ChannelDetails, error)
	FindByUserID(userID uuid.UUID) ([]*ChannelDetails, error)
	FindByChannelType(channelType user.ChannelType) ([]*ChannelDetails, error)
	FindByStatus(status bool) ([]*ChannelDetails, error)
	Update(details *ChannelDetails) error
	Delete(id uuid.UUID) error
}
