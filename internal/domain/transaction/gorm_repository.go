package transaction

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormTransactionRepository struct {
	DB *gorm.DB
}

func NewGormTransactionRepository(db *gorm.DB) *GormTransactionRepository {
	return &GormTransactionRepository{DB: db}
}

func (r *GormTransactionRepository) Create(tx *Transaction) error {
	return r.DB.Create(tx).Error
}

func (r *GormTransactionRepository) FindByID(id uuid.UUID) (*Transaction, error) {
	var tx Transaction
	if err := r.DB.First(&tx, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &tx, nil
}

func (r *GormTransactionRepository) FindByUserID(userID uuid.UUID) ([]*Transaction, error) {
	var txs []*Transaction
	// Join with profiles table to filter by user_id
	if err := r.DB.Preload("User").Preload("ChannelDetails").Where("user_id = ?", userID).Find(&txs).Error; err != nil {
		return nil, err
	}
	return txs, nil
}

func (r *GormTransactionRepository) Update(tx *Transaction) error {
	return r.DB.Save(tx).Error
}

func (r *GormTransactionRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&Transaction{}, "id = ?", id).Error
}
