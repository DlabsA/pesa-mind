package analytics

import "github.com/google/uuid"

type Repository interface {
	Create(s *AnalyticsSnapshot) error
	FindByID(id uuid.UUID) (*AnalyticsSnapshot, error)
	FindByUserID(userID uuid.UUID, limit, offset int) ([]AnalyticsSnapshot, error)
	Delete(id uuid.UUID) error
}
