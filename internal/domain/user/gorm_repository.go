package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

type UserProfile struct {
	user     *User
	username string
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{DB: db}
}

func (r *GormUserRepository) Create(userProfile UserProfile) error {
	err := r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(userProfile.user).Error; err != nil {
			return err
		}
		profile := &Profile{
			UserID:   userProfile.user.ID,
			Username: userProfile.username, // Default username as email
			Type:     "Free",               // Default profile type
			Balance:  0.0,                  // Default balance
		}
		if err := tx.Create(profile).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (r *GormUserRepository) FindByID(id uuid.UUID) (*User, *Profile, error) {
	var user User
	var profile Profile
	if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, nil, err
	}
	if err := r.DB.First(&profile, "user_id = ?", user.ID).Error; err != nil {
		return &user, nil, err
	}
	return &user, &profile, nil
}

func (r *GormUserRepository) FindByEmail(email string) (*User, *Profile, error) {
	var user User
	var profile Profile
	if err := r.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, nil, err
	}
	if err := r.DB.First(&profile, "user_id = ?", user.ID).Error; err != nil {
		return &user, nil, err
	}
	return &user, &profile, nil
}

func (r *GormUserRepository) Update(user *User) error {
	return r.DB.Save(user).Error
}

func (r *GormUserRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&User{}, "id = ?", id).Error
}
