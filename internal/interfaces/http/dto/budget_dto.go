package dto

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
