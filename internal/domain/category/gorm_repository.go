package category

import (
	"pesa-mind/internal/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormChannelDetailsRepository struct {
	DB *gorm.DB
}

func NewGormChannelDetailsRepository(db *gorm.DB) *GormChannelDetailsRepository {
	return &GormChannelDetailsRepository{DB: db}
}

func (r *GormChannelDetailsRepository) Create(channelDetails *ChannelDetails) error {
	return r.DB.Create(channelDetails).Error
}

func (r *GormChannelDetailsRepository) FindByID(id uuid.UUID) (*ChannelDetails, error) {
	var channelDetails ChannelDetails
	if err := r.DB.Where("id = ?", id).Find(&channelDetails).Error; err != nil {
		return nil, err
	}
	return &channelDetails, nil
}

func (r *GormChannelDetailsRepository) FindByUserID(userID uuid.UUID) ([]*ChannelDetails, error) {
	var channelDetails []*ChannelDetails
	if err := r.DB.Where("user_id = ?", userID).Find(&channelDetails).Error; err != nil {
		return nil, err
	}
	return channelDetails, nil
}

func (r *GormChannelDetailsRepository) FindByChannelType(channelType user.ChannelType) ([]*ChannelDetails, error) {
	var channelDetails []*ChannelDetails
	if err := r.DB.Where("channel_type = ?", channelType).Find(&channelDetails).Error; err != nil {
		return nil, err
	}
	return channelDetails, nil
}

func (r *GormChannelDetailsRepository) FindByStatus(status bool) ([]*ChannelDetails, error) {
	var channelDetails []*ChannelDetails
	if err := r.DB.Where("active = ?", status).Find(&channelDetails).Error; err != nil {
		return nil, err
	}
	return channelDetails, nil
}

func (r *GormChannelDetailsRepository) Update(channelDetails *ChannelDetails) error {
	return r.DB.Save(channelDetails).Error
}

func (r *GormChannelDetailsRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&ChannelDetails{}, "id = ?", id).Error
}
