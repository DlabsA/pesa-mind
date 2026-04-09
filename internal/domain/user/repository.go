package user

import "github.com/google/uuid"

type UserRepository interface {
	Create(profile UserProfile) error
	FindByID(id uuid.UUID) (*User, *Profile, error)
	FindByEmail(email string) (*User, *Profile, error)
	Update(user *User) error
	Delete(id uuid.UUID) error
}

type ProfileRepository interface {
	Create(profile *Profile) error
	FindByID(id uuid.UUID) (*Profile, error)
	Update(profile *Profile) error
	Delete(id uuid.UUID) error
}
