package dto

type (
	CreateCategoryRequest struct {
		Name     string  `json:"name" binding:"required"`
		Type     string  `json:"type" binding:"required"`
		ParentID *string `json:"parent_id"`
	}

	CategoryResponse struct {
		ID       string  `json:"id"`
		UserID   string  `json:"user_id"`
		Name     string  `json:"name"`
		Type     string  `json:"type"`
		ParentID *string `json:"parent_id,omitempty"`
	}

	CreateChannelDetailsRequest struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
		ChannelType string `json:"channel_type" binding:"required"`
		ChannelDesc string `json:"channel_desc" binding:"required"`
		Status      bool   `json:"status" binding:"required"`
	}

	UpdateChannelDetailsRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Status      *bool  `json:"status"`
	}

	ChannelDetailsResponse struct {
		ID          string `json:"id"`
		UserID      string `json:"user_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ChannelType string `json:"channel_type"`
		ChannelDesc string `json:"channel_desc"`
		Status      bool   `json:"status"`
	}
)
