package budget

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

func (r *GormRepository) Create(b *Budget) error {
	return r.db.Create(b).Error
}

func (r *GormRepository) FindByID(id uuid.UUID) (*Budget, error) {
	var b Budget
	if err := r.db.First(&b, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *GormRepository) FindByUserID(userID uuid.UUID, limit, offset int) ([]Budget, error) {
	var budgets []Budget
	err := r.db.Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&budgets).Error
	return budgets, err
}

func (r *GormRepository) Update(b *Budget) error {
	return r.db.Save(b).Error
}

func (r *GormRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&Budget{}, "id = ?", id).Error
}
