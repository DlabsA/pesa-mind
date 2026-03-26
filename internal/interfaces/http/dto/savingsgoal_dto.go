package dto

type CreateSavingsGoalRequest struct {
	Title    string  `json:"title" binding:"required"`
	Target   float64 `json:"target" binding:"required"`
	Deadline *int64  `json:"deadline"`
	AutoSave bool    `json:"auto_save"`
}

type SavingsGoalResponse struct {
	ID       string  `json:"id"`
	UserID   string  `json:"user_id"`
	Title    string  `json:"title"`
	Target   float64 `json:"target"`
	Current  float64 `json:"current"`
	Deadline *int64  `json:"deadline,omitempty"`
	AutoSave bool    `json:"auto_save"`
}
