package automation

import "github.com/google/uuid"

// Repository interface for SMSAutomation
// Allows for mocking and easy testing

type Repository interface {
	Create(rule *SMSAutomation) error
	FindByUserID(userID uuid.UUID) ([]SMSAutomation, error)
	FindByID(id uuid.UUID) (*SMSAutomation, error)
	Update(rule *SMSAutomation) error
	Delete(id uuid.UUID) error
}
