package user

import (
	"errors"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo UserRepository
}

var ErrInvalidCurrentPassword = errors.New("invalid current password")

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

// Update updates user's email and/or profile username. It returns the updated
// User and Profile. If no user is found, (nil, nil, nil) is returned.
func (s *Service) Update(id uuid.UUID, email, username string) (*User, *Profile, error) {
	u, p, err := s.repo.FindByID(id)
	if err != nil {
		return nil, nil, err
	}
	if u == nil {
		return nil, nil, nil
	}

	if email != "" {
		u.Email = email
	}
	if username != "" {
		if p == nil {
			// create a minimal profile if none exists
			p = &Profile{
				UserID:   u.ID,
				Username: username,
				Type:     "Free",
				Balance:  0.0,
			}
			u.Profile = p
		} else {
			p.Username = username
			u.Profile = p
		}
	}

	// Persist changes. The repository's Update saves the user; ensure profile
	// is attached so GORM can save associations if configured.
	if err := s.repo.Update(u); err != nil {
		return nil, nil, err
	}

	return u, p, nil
}

func (s *Service) ChangePassword(id uuid.UUID, currentPassword, newPassword string) (*User, error) {
	u, _, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		// follow existing service.Update pattern of returning nils when user not found
		return nil, nil
	}

	// Verify current password
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(currentPassword)); err != nil {
		// wrong current password
		return nil, ErrInvalidCurrentPassword
	}

	// Hash new password
	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Update user model and persist
	u.PasswordHash = string(hashed)
	if err := s.repo.Update(u); err != nil {
		return nil, err
	}

	return u, nil
}
