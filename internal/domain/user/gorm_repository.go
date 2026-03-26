package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{DB: db}
}

func (r *GormUserRepository) Create(user *User) error {
	return r.DB.Create(user).Error
}

func (r *GormUserRepository) FindByID(id uuid.UUID) (*User, error) {
	var user User
	if err := r.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) FindByEmail(email string) (*User, error) {
	var user User
	if err := r.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormUserRepository) Update(user *User) error {
	return r.DB.Save(user).Error
}

func (r *GormUserRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&User{}, "id = ?", id).Error
}
