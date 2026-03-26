package notification

import "github.com/google/uuid"

type Repository interface {
	Create(notification *Notification) error
	ListByUser(userID uuid.UUID, limit, offset int) ([]Notification, error)
	MarkAsRead(id uuid.UUID) error
	FindByID(id uuid.UUID) (*Notification, error)
}
