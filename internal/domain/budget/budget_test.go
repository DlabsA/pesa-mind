package budget

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockRepo struct {
	budgets []Budget
}

func (m *mockRepo) Create(b *Budget) error {
	m.budgets = append(m.budgets, *b)
	return nil
}
func (m *mockRepo) FindByID(id uuid.UUID) (*Budget, error) {
	for _, b := range m.budgets {
		if b.ID == id {
			return &b, nil
		}
	}
	return nil, nil
}
func (m *mockRepo) FindByUserID(userID uuid.UUID, limit, offset int) ([]Budget, error) {
	var res []Budget
	for _, b := range m.budgets {
		if b.UserID == userID {
			res = append(res, b)
		}
	}
	return res, nil
}
func (m *mockRepo) Update(b *Budget) error    { return nil }
func (m *mockRepo) Delete(id uuid.UUID) error { return nil }

func TestService_CreateAndList(t *testing.T) {
	repo := &mockRepo{}
	svc := NewService(repo)
	userID := uuid.New()
	b, err := svc.Create(userID, "Test Budget", 1000, "monthly", 1700000000, 1702592000)
	assert.NoError(t, err)
	assert.Equal(t, "Test Budget", b.Name)
	budgets, err := svc.List(userID, 10, 0)
	assert.NoError(t, err)
	assert.Len(t, budgets, 1)
}
