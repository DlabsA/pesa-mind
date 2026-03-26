package category

import "github.com/google/uuid"

type CategoryRepository interface {
	Create(category *Category) error
	FindByID(id uuid.UUID) (*Category, error)
	FindByUserID(userID uuid.UUID) ([]*Category, error)
	Update(category *Category) error
	Delete(id uuid.UUID) error
}
