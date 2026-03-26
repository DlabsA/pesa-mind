package account

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormAccountRepository struct {
	DB *gorm.DB
}

func NewGormAccountRepository(db *gorm.DB) *GormAccountRepository {
	return &GormAccountRepository{DB: db}
}

func (r *GormAccountRepository) Create(account *Account) error {
	return r.DB.Create(account).Error
}

func (r *GormAccountRepository) FindByID(id uuid.UUID) (*Account, error) {
	var account Account
	if err := r.DB.First(&account, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *GormAccountRepository) FindByUserID(userID uuid.UUID) ([]*Account, error) {
	var accounts []*Account
	if err := r.DB.Where("user_id = ?", userID).Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *GormAccountRepository) Update(account *Account) error {
	return r.DB.Save(account).Error
}

func (r *GormAccountRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&Account{}, "id = ?", id).Error
}
