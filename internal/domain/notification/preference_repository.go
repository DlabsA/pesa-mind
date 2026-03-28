package notification

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PreferenceRepository interface {
	GetByUserID(userID uuid.UUID) (*Preference, error)
	Set(userID uuid.UUID, inApp, push, email bool) error
}

type GormPreferenceRepository struct {
	DB *gorm.DB
}

func NewGormPreferenceRepository(db *gorm.DB) *GormPreferenceRepository {
	return &GormPreferenceRepository{DB: db}
}

func (r *GormPreferenceRepository) GetByUserID(userID uuid.UUID) (*Preference, error) {
	var pref Preference
	err := r.DB.Where("user_id = ?", userID).First(&pref).Error
	if err != nil {
		return nil, err
	}
	return &pref, nil
}

func (r *GormPreferenceRepository) Set(userID uuid.UUID, inApp, push, email bool) error {
	var pref Preference
	err := r.DB.Where("user_id = ?", userID).First(&pref).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	pref.UserID = userID
	pref.InApp = inApp
	pref.Push = push
	pref.Email = email
	if err == gorm.ErrRecordNotFound {
		return r.DB.Create(&pref).Error
	}
	return r.DB.Save(&pref).Error
}
