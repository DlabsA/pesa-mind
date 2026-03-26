package automation

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{DB: db}
}

func (r *GormRepository) Create(rule *SMSAutomation) error {
	return r.DB.Create(rule).Error
}

func (r *GormRepository) FindByUserID(userID uuid.UUID) ([]SMSAutomation, error) {
	var rules []SMSAutomation
	err := r.DB.Where("user_id = ? AND deleted_at IS NULL", userID).Find(&rules).Error
	return rules, err
}

func (r *GormRepository) FindByID(id uuid.UUID) (*SMSAutomation, error) {
	var rule SMSAutomation
	err := r.DB.Where("id = ? AND deleted_at IS NULL", id).First(&rule).Error
	if err != nil {
		return nil, err
	}
	return &rule, nil
}

func (r *GormRepository) Update(rule *SMSAutomation) error {
	return r.DB.Save(rule).Error
}

func (r *GormRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&SMSAutomation{}, "id = ?", id).Error
}
