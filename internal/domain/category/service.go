package category

import "github.com/google/uuid"

type Service struct {
	repo CategoryRepository
}

func NewService(repo CategoryRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(userID uuid.UUID, name, catType string, parentID *uuid.UUID) (*Category, error) {
	category := &Category{
		ID:       uuid.New(),
		UserID:   userID,
		Name:     name,
		Type:     catType,
		ParentID: parentID,
	}
	if err := s.repo.Create(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *Service) GetByID(id uuid.UUID) (*Category, error) {
	return s.repo.FindByID(id)
}

func (s *Service) GetByUserID(userID uuid.UUID) ([]*Category, error) {
	return s.repo.FindByUserID(userID)
}
