package user

import (
	"github.com/google/uuid"
)

type Service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{repo: repo}
}

// ...existing code...

func (s *Service) Register(email, passwordHash string, username string) (*User, error) {
	// Default username to email if not provided
	if username == "" {
		username = email
	}

	user := &User{
		Email:        email,
		PasswordHash: passwordHash,
	}
	userProfile := UserProfile{
		user:     user,
		username: username,
	}
	if err := s.repo.Create(userProfile); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetByID(id uuid.UUID) (*User, *Profile, error) {
	return s.repo.FindByID(id)
}

func (s *Service) GetByEmail(email string) (*User, *Profile, error) {
	return s.repo.FindByEmail(email)
}
