package savingsgoal

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Create(goal *SavingsGoal) error {
	return r.db.Create(goal).Error
}

func (r *GormRepository) FindByID(id uuid.UUID) (*SavingsGoal, error) {
	var g SavingsGoal
	if err := r.db.First(&g, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *GormRepository) FindByUserID(userID uuid.UUID, limit, offset int) ([]SavingsGoal, error) {
	var goals []SavingsGoal
	err := r.db.Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&goals).Error
	return goals, err
}

func (r *GormRepository) Update(goal *SavingsGoal) error {
	return r.db.Save(goal).Error
}

func (r *GormRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&SavingsGoal{}, "id = ?", id).Error
}
