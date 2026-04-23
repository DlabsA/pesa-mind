package dto

import "github.com/google/uuid"

type CreateBudgetRequest struct {
	Name      string  `json:"name" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
	Period    string  `json:"period" binding:"required"`
	StartDate int64   `json:"start_date" binding:"required"`
	EndDate   int64   `json:"end_date" binding:"required"`
}

type BudgetResponse struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	Name      string  `json:"name"`
	Amount    float64 `json:"amount"`
	Period    string  `json:"period"`
	StartDate int64   `json:"start_date"`
	EndDate   int64   `json:"end_date"`
}

// BudgetTransaction DTO
type BudgetTransactionRequest struct {
	Name   string  `json:"name" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
	Type   string  `json:"type" binding:"required,oneof=income expense saving"`
}

type BudgetTransactionResponse struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
	CreatedAt int64   `json:"created_at"`
}

// MonthlyBudget DTOs
type CreateMonthlyBudgetRequest struct {
	YearlyBudgetID uuid.UUID                  `json:"yearly_budget_id" binding:"required"`
	Month          int                        `json:"month" binding:"required,min=1,max=12"`
	Year           int64                      `json:"year" binding:"required"`
	Transactions   []BudgetTransactionRequest `json:"transactions"`
}

type MonthlyBudgetResponse struct {
	ID                string                      `json:"id"`
	UserID            string                      `json:"user_id"`
	YearlyBudgetID    string                      `json:"yearly_budget_id"`
	Month             int                         `json:"month"`
	Year              int64                       `json:"year"`
	TotalExpenditures uint64                      `json:"total_expenditures"`
	TotalIncome       uint64                      `json:"total_income"`
	TotalSavings      uint64                      `json:"total_savings"`
	TotalTransactions uint64                      `json:"total_transactions"`
	Transactions      []BudgetTransactionResponse `json:"transactions,omitempty"`
	CreatedAt         int64                       `json:"created_at"`
	UpdatedAt         int64                       `json:"updated_at"`
}

// YearlyBudget DTOs
type CreateYearlyBudgetRequest struct {
	Year         int64                      `json:"year" binding:"required"`
	Transactions []BudgetTransactionRequest `json:"transactions"`
}

type YearlyBudgetResponse struct {
	ID                string                      `json:"id"`
	UserID            string                      `json:"user_id"`
	Year              int64                       `json:"year"`
	TotalExpenditures uint64                      `json:"total_expenditures"`
	TotalIncome       uint64                      `json:"total_income"`
	TotalSavings      uint64                      `json:"total_savings"`
	TotalTransactions uint64                      `json:"total_transactions"`
	Transactions      []BudgetTransactionResponse `json:"transactions,omitempty"`
	CreatedAt         int64                       `json:"created_at"`
	UpdatedAt         int64                       `json:"updated_at"`
}

type UpdateBudgetRequest struct {
	Name      string  `json:"name"`
	Amount    float64 `json:"amount"`
	Period    string  `json:"period"`
	StartDate int64   `json:"start_date"`
	EndDate   int64   `json:"end_date"`
}
