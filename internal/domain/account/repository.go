package account

import "github.com/google/uuid"

type AccountRepository interface {
	Create(account *Account) error
	FindByID(id uuid.UUID) (*Account, error)
	FindByUserID(userID uuid.UUID) ([]*Account, error)
	Update(account *Account) error
	Delete(id uuid.UUID) error
}
