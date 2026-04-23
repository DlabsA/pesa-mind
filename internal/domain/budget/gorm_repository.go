package budget

import (
	"time"

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

func (r *GormRepository) CreateMonthlyBudget(mb *MonthlyBudget) error {
	return r.db.Create(mb).Error
}

func (r *GormRepository) FindMonthlyBudgetByID(ID uuid.UUID) (*MonthlyBudget, error) {
	var mb MonthlyBudget
	if err := r.db.First(&mb, "id = ?", ID).Error; err != nil {
		return nil, err
	}
	return &mb, nil
}

func (r *GormRepository) FindMonthlyBudgetsByYearlyBudgetID(yearlyBudgetID uuid.UUID) ([]MonthlyBudget, error) {
	var mbs []MonthlyBudget
	err := r.db.Where("yearly_budget_id = ?", yearlyBudgetID).Find(&mbs).Error
	return mbs, err
}

func (r *GormRepository) FindMonthlyBudgetByUserIDAndMonthYear(userID uuid.UUID, month time.Month, year int64) (*MonthlyBudget, error) {
	var mb MonthlyBudget
	if err := r.db.Where("user_id = ? AND month = ? AND year = ?", userID, month, year).First(&mb).Error; err != nil {
		return nil, err
	}
	return &mb, nil
}

func (r *GormRepository) FindMonthlyBudgetsByUserID(userID uuid.UUID, limit, offset int) ([]MonthlyBudget, error) {
	var mbs []MonthlyBudget
	err := r.db.Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&mbs).Error
	return mbs, err
}

func (r *GormRepository) UpdateMonthlyBudget(mb *MonthlyBudget) error {
	return r.db.Save(mb).Error
}

func (r *GormRepository) DeleteMonthlyBudget(id uuid.UUID) error {
	return r.db.Delete(&MonthlyBudget{}, "id = ?", id).Error
}

func (r *GormRepository) CreateYearlyBudget(yb *YearlyBudget) error {
	return r.db.Create(yb).Error
}

func (r *GormRepository) FindYearlyBudgetByID(id uuid.UUID) (*YearlyBudget, error) {
	var yb YearlyBudget
	if err := r.db.First(&yb, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &yb, nil
}

func (r *GormRepository) FindYearlyBudgetsByUserID(userID uuid.UUID, limit, offset int) ([]YearlyBudget, error) {
	var ybs []YearlyBudget
	err := r.db.Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&ybs).Error
	return ybs, err
}

func (r *GormRepository) UpdateYearlyBudget(yb *YearlyBudget) error {
	return r.db.Save(yb).Error
}

func (r *GormRepository) DeleteYearlyBudget(id uuid.UUID) error {
	return r.db.Delete(&YearlyBudget{}, "id = ?", id).Error
}
