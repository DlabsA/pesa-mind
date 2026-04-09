package dto

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Username string `json:"username" binding:"omitempty,min=3,max=50"` // Optional, defaults to email
}

type UserResponse struct {
	ID      string       `json:"id"`
	Email   string       `json:"email"`
	Profile *ProfileData `json:"profile,omitempty"`
}

type ProfileData struct {
	ID       string  `json:"id"`
	UserID   string  `json:"user_id"`
	Username string  `json:"username"`
	Type     string  `json:"type"`
	Balance  float64 `json:"balance"`
}
