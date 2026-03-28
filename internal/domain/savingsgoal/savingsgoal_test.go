package savingsgoal

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type mockRepo struct {
	goals []SavingsGoal
}

func (m *mockRepo) Create(goal *SavingsGoal) error {
	m.goals = append(m.goals, *goal)
	return nil
}
func (m *mockRepo) FindByID(id uuid.UUID) (*SavingsGoal, error) {
	for _, g := range m.goals {
		if g.ID == id {
			return &g, nil
		}
	}
	return nil, nil
}
func (m *mockRepo) FindByUserID(userID uuid.UUID, limit, offset int) ([]SavingsGoal, error) {
	var res []SavingsGoal
	for _, g := range m.goals {
		if g.UserID == userID {
			res = append(res, g)
		}
	}
	return res, nil
}
func (m *mockRepo) Update(goal *SavingsGoal) error { return nil }
func (m *mockRepo) Delete(id uuid.UUID) error      { return nil }

func TestService_CreateAndList(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo, nil)
	userID := uuid.New()
	goal, err := svc.Create(userID, "Trip to Japan", 5000, nil, true)
	assert.NoError(t, err)
	assert.Equal(t, "Trip to Japan", goal.Title)
	goals, err := svc.List(userID, 10, 0)
	assert.NoError(t, err)
	assert.Len(t, goals, 1)
}
