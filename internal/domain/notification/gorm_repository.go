package notification

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

func (r *GormRepository) Create(notification *Notification) error {
	return r.DB.Create(notification).Error
}

func (r *GormRepository) ListByUser(userID uuid.UUID, limit, offset int) ([]Notification, error) {
	var notifications []Notification
	err := r.DB.Where("user_id = ? AND deleted_at IS NULL", userID).
		Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&notifications).Error
	return notifications, err
}

func (r *GormRepository) MarkAsRead(id uuid.UUID) error {
	return r.DB.Model(&Notification{}).Where("id = ?", id).Update("status", "read").Error
}

func (r *GormRepository) FindByID(id uuid.UUID) (*Notification, error) {
	var n Notification
	err := r.DB.Where("id = ? AND deleted_at IS NULL", id).First(&n).Error
	if err != nil {
		return nil, err
	}
	return &n, nil
}
