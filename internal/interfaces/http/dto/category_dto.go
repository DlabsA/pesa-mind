package dto

type CreateCategoryRequest struct {
	Name     string  `json:"name" binding:"required"`
	Type     string  `json:"type" binding:"required"`
	ParentID *string `json:"parent_id"`
}

type CategoryResponse struct {
	ID       string  `json:"id"`
	UserID   string  `json:"user_id"`
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	ParentID *string `json:"parent_id,omitempty"`
}

type CreateChannelDetailsRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ChannelType string `json:"channel_type" binding:"required"`
	Status      bool   `json:"status" binding:"required"`
}

type UpdateChannelDetailsRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      *bool  `json:"status"`
}

type ChannelDetailsResponse struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ChannelType string `json:"channel_type"`
	Status      bool   `json:"status"`
}
