package notification

import "github.com/google/uuid"

type PreferenceService struct {
	repo PreferenceRepository
}

func NewPreferenceService(repo PreferenceRepository) *PreferenceService {
	return &PreferenceService{repo: repo}
}

func (s *PreferenceService) Get(userID uuid.UUID) (*Preference, error) {
	return s.repo.GetByUserID(userID)
}

func (s *PreferenceService) Set(userID uuid.UUID, inApp, push, email bool) error {
	return s.repo.Set(userID, inApp, push, email)
}
