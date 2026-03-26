package analytics

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

func (r *GormRepository) Create(s *AnalyticsSnapshot) error {
	return r.db.Create(s).Error
}

func (r *GormRepository) FindByID(id uuid.UUID) (*AnalyticsSnapshot, error) {
	var snap AnalyticsSnapshot
	if err := r.db.First(&snap, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &snap, nil
}

func (r *GormRepository) FindByUserID(userID uuid.UUID, limit, offset int) ([]AnalyticsSnapshot, error) {
	var snaps []AnalyticsSnapshot
	err := r.db.Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&snaps).Error
	return snaps, err
}

func (r *GormRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&AnalyticsSnapshot{}, "id = ?", id).Error
}
