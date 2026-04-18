package dto

type CreateTransactionRequest struct {
	ProfileID  string  `json:"profile_id" binding:"required,uuid4"`
	CategoryID string  `json:"category_id" binding:"required,uuid4"`
	Amount     float64 `json:"amount" binding:"required"`
	Type       string  `json:"type" binding:"required"`
	Note       string  `json:"note"`
	Date       int64   `json:"date" binding:"required"` // Unix timestamp
	UserID     string  `json:"user_id" binding:"required,uuid4"`
}

type TransactionResponse struct {
	ID         string  `json:"id"`
	UserID     string  `json:"user_id"`
	ProfileID  string  `json:"profile_id"`
	CategoryID string  `json:"category_id"`
	Amount     float64 `json:"amount"`
	Type       string  `json:"type"`
	Note       string  `json:"note"`
	Date       int64   `json:"date"`
}
