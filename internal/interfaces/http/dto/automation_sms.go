package dto

type AutomationSMSRequest struct {
	UserID      string  `json:"user_id" binding:"required,uuid4"`
	Amount      float64 `json:"amount" binding:"required"`
	CategoryID  string  `json:"category_id" binding:"required,uuid4"`
	Description string  `json:"description" binding:"required"`
	OccurredAt  string  `json:"occurred_at" binding:"required,datetime"`
}

type AutomationSMSResponse struct {
	TransactionID string `json:"transaction_id"`
}
