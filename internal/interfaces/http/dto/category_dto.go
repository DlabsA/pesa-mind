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
