package budget

import "github.com/google/uuid"

type Repository interface {
	Create(b *Budget) error
	FindByID(id uuid.UUID) (*Budget, error)
	FindByUserID(userID uuid.UUID, limit, offset int) ([]Budget, error)
	Update(b *Budget) error
	Delete(id uuid.UUID) error
}
