package analytics

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type mockRepo struct {
	snaps []AnalyticsSnapshot
}

func (m *mockRepo) Create(s *AnalyticsSnapshot) error {
	m.snaps = append(m.snaps, *s)
	return nil
}
func (m *mockRepo) FindByID(id uuid.UUID) (*AnalyticsSnapshot, error) {
	for _, s := range m.snaps {
		if s.ID == id {
			return &s, nil
		}
	}
	return nil, nil
}
func (m *mockRepo) FindByUserID(userID uuid.UUID, limit, offset int) ([]AnalyticsSnapshot, error) {
	var res []AnalyticsSnapshot
	for _, s := range m.snaps {
		if s.UserID == userID {
			res = append(res, s)
		}
	}
	return res, nil
}
func (m *mockRepo) Delete(id uuid.UUID) error { return nil }

func TestService_CreateAndList(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()
	snap, err := svc.Create(userID, "trend", `{"spending":100}`, "monthly")
	assert.NoError(t, err)
	assert.Equal(t, "trend", snap.Type)
	snaps, err := svc.List(userID, 10, 0)
	assert.NoError(t, err)
	assert.Len(t, snaps, 1)
}
