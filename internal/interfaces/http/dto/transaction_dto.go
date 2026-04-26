package dto

type CreateTransactionRequest struct {
	ChannelDetailsID string  `json:"channel_details_id" binding:"required,uuid4"`
	Amount           float64 `json:"amount" binding:"required"`
	Type             string  `json:"type" binding:"required"`
	Note             string  `json:"note"`
	UserID           string  `json:"user_id" binding:"required,uuid4"`
}

type TransactionResponse struct {
	ID                 string  `json:"id"`
	Username           string  `json:"username"`
	ChannelDetailsName string  `json:"channel_details_name"`
	Amount             float64 `json:"amount"`
	Type               string  `json:"type"`
	Note               string  `json:"note"`
}
