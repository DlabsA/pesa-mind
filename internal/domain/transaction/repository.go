package transaction

import "github.com/google/uuid"

type TransactionRepository interface {
	Create(tx *Transaction) error
	FindByID(id uuid.UUID) (*Transaction, error)
	FindByUserID(userID uuid.UUID) ([]*Transaction, error)
	FindByUserIDAndType(userID uuid.UUID, txType string) ([]*Transaction, error)
	Update(tx *Transaction) error
	Delete(id uuid.UUID) error
}
