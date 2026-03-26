package notification

import (
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Send(userID uuid.UUID, notifType, title, content string) error {
	n := &Notification{
		UserID:  userID,
		Type:    notifType,
		Title:   title,
		Content: content,
		Status:  "unread",
	}
	return s.repo.Create(n)
}

// SendPush is a stub for future push notification integration.
func (s *Service) SendPush(userID uuid.UUID, title, content string) error {
	// TODO: Integrate with push notification provider
	return s.Send(userID, "push", title, content)
}

// SendEmail is a stub for future email notification integration.
func (s *Service) SendEmail(userID uuid.UUID, title, content string) error {
	// TODO: Integrate with email provider
	return s.Send(userID, "email", title, content)
}

func (s *Service) List(userID uuid.UUID, limit, offset int) ([]Notification, error) {
	return s.repo.ListByUser(userID, limit, offset)
}

func (s *Service) MarkAsRead(id uuid.UUID) error {
	return s.repo.MarkAsRead(id)
}
