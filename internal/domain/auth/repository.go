package auth

import (
	"github.com/google/uuid"
	"time"
)

type RefreshTokenRepository interface {
	Create(token *RefreshToken) error
	FindByID(id uuid.UUID) (*RefreshToken, error)
	FindByUserID(userID uuid.UUID) ([]*RefreshToken, error)
	Revoke(id uuid.UUID) error
	DeleteExpired(now time.Time) error
}
