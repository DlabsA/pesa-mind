package category

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormCategoryRepository struct {
	DB *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) *GormCategoryRepository {
	return &GormCategoryRepository{DB: db}
}

func (r *GormCategoryRepository) Create(category *Category) error {
	return r.DB.Create(category).Error
}

func (r *GormCategoryRepository) FindByID(id uuid.UUID) (*Category, error) {
	var category Category
	if err := r.DB.First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *GormCategoryRepository) FindByUserID(userID uuid.UUID) ([]*Category, error) {
	var categories []*Category
	if err := r.DB.Where("user_id = ?", userID).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *GormCategoryRepository) Update(category *Category) error {
	return r.DB.Save(category).Error
}

func (r *GormCategoryRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&Category{}, "id = ?", id).Error
}
