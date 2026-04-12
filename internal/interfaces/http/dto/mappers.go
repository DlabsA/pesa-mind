package dto

import (
	"pesa-mind/internal/domain/user"
)

func ToProfileDTO(p *user.Profile) *ProfileData {
	if p == nil {
		return nil
	}
	return &ProfileData{
		ID:       p.ID.String(),
		UserID:   p.UserID.String(),
		Username: p.Username,
		Type:     p.Type,
		Balance:  p.Balance,
	}
}

// ToUserResponse maps domain User and Profile to API UserResponse. If profile
// is nil, the user's embedded profile will be used.
func ToUserResponse(u *user.User, profile *user.Profile) *UserResponse {
	if u == nil {
		return nil
	}
	selected := profile
	if selected == nil {
		selected = u.Profile
	}
	return &UserResponse{
		ID:      u.ID.String(),
		Email:   u.Email,
		Profile: ToProfileDTO(selected),
	}
}

