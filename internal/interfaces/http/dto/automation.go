package dto

type CreateSMSAutomationRequest struct {
	Pattern string `json:"pattern" binding:"required"`
}

type SMSAutomationResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Pattern   string `json:"pattern"`
	Enabled   bool   `json:"enabled"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
