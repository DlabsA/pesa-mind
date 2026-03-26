package dto

type CreateAnalyticsSnapshotRequest struct {
	Type   string `json:"type" binding:"required"`
	Data   string `json:"data" binding:"required"`
	Period string `json:"period" binding:"required"`
}

type AnalyticsSnapshotResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Type      string `json:"type"`
	Data      string `json:"data"`
	Period    string `json:"period"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
