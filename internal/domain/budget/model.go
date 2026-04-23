package budget

import (
	"pesa-mind/internal/domain/user"
	"pesa-mind/internal/infrastructure/utils"
	"time"

	"github.com/google/uuid"
)

type Budget struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Name      string     `gorm:"not null" json:"name"`
	Amount    float64    `gorm:"not null" json:"amount"`
	Period    string     `gorm:"not null" json:"period"` // e.g. monthly, weekly
	StartDate time.Time  `gorm:"not null" json:"start_date"`
	EndDate   time.Time  `gorm:"not null" json:"end_date"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type BudgetTransaction struct {
	utils.BaseModel
	MonthlyBudgetID uuid.UUID             `gorm:"type:uuid;index" json:"monthly_budget_id"`
	YearlyBudgetID  uuid.UUID             `gorm:"type:uuid;index" json:"yearly_budget_id"`
	Name            string                `gorm:"not null" json:"name"`
	Amount          float64               `gorm:"not null" json:"amount"`
	Type            utils.TransactionType `gorm:"not null" json:"type"`
}

type MonthlyBudget struct {
	utils.BaseModel
	UserID             uuid.UUID           `gorm:"type:uuid;not null;index" json:"user_id"`
	User               *user.User          `gorm:"foreignKey:UserID" json:"user"`
	YearlyBudgetID     uuid.UUID           `gorm:"type:uuid;not null;index;uniqueIndex:idx_yearly_month" json:"yearly_budget_id"`
	YearlyBudget       *YearlyBudget       `gorm:"foreignKey:YearlyBudgetID" json:"yearly_budget"`
	Month              time.Month          `gorm:"not null;uniqueIndex:idx_yearly_month" json:"month"`
	Year               int64               `gorm:"not null" json:"year"`
	BudgetTransactions []BudgetTransaction `gorm:"foreignKey:MonthlyBudgetID" json:"budget_transactions"`
	TotalExpenditures  uint64              `gorm:"not null" json:"total_expenditures"`
	TotalIncome        uint64              `gorm:"null" json:"total_income"`
	TotalSavings       uint64              `gorm:"null" json:"total_savings"`
	TotalTransactions  uint64              `gorm:"null" json:"total_transactions"`
}

type YearlyBudget struct {
	utils.BaseModel
	UserID             uuid.UUID           `gorm:"type:uuid;not null;index" json:"user_id"`
	User               *user.User          `gorm:"foreignKey:UserID" json:"user"`
	Year               int64               `gorm:"not null;uniqueIndex:idx_user_year" json:"year"`
	MonthlyBudget      []MonthlyBudget     `gorm:"foreignKey:YearlyBudgetID" json:"monthly_budgets"`
	BudgetTransactions []BudgetTransaction `gorm:"foreignKey:YearlyBudgetID" json:"budget_transactions"`
	TotalExpenditures  uint64              `gorm:"not null" json:"total_expenditures"`
	TotalIncome        uint64              `gorm:"null" json:"total_income"`
	TotalSavings       uint64              `gorm:"null" json:"total_savings"`
	TotalTransactions  uint64              `gorm:"null" json:"total_transactions"`
}
