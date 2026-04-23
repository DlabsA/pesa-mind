package budget

import (
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	Create(b *Budget) error
	FindByID(id uuid.UUID) (*Budget, error)
	FindByUserID(userID uuid.UUID, limit, offset int) ([]Budget, error)
	Update(b *Budget) error
	Delete(id uuid.UUID) error
	CreateMonthlyBudget(mb *MonthlyBudget) error
	FindMonthlyBudgetByID(id uuid.UUID) (*MonthlyBudget, error)
	FindMonthlyBudgetsByYearlyBudgetID(yearlyBudgetID uuid.UUID) ([]MonthlyBudget, error)
	FindMonthlyBudgetByUserIDAndMonthYear(userID uuid.UUID, month time.Month, year int64) (*MonthlyBudget, error)
	FindMonthlyBudgetsByUserID(userID uuid.UUID, limit, offset int) ([]MonthlyBudget, error)
	UpdateMonthlyBudget(mb *MonthlyBudget) error
	DeleteMonthlyBudget(id uuid.UUID) error
	CreateYearlyBudget(yb *YearlyBudget) error
	FindYearlyBudgetByID(id uuid.UUID) (*YearlyBudget, error)
	FindYearlyBudgetsByUserID(userID uuid.UUID, limit, offset int) ([]YearlyBudget, error)
	UpdateYearlyBudget(yb *YearlyBudget) error
	DeleteYearlyBudget(id uuid.UUID) error
}
