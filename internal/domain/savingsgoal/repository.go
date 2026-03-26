package savingsgoal

import "github.com/google/uuid"

type Repository interface {
	Create(goal *SavingsGoal) error
	FindByID(id uuid.UUID) (*SavingsGoal, error)
	FindByUserID(userID uuid.UUID, limit, offset int) ([]SavingsGoal, error)
	Update(goal *SavingsGoal) error
	Delete(id uuid.UUID) error
}
