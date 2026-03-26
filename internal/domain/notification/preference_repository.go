package notification

import "github.com/google/uuid"

type PreferenceRepository interface {
	GetByUserID(userID uuid.UUID) (*Preference, error)
	Set(userID uuid.UUID, inApp, push, email bool) error
}

type GormPreferenceRepository struct {
	DB DBProvider
}

type DBProvider interface {
	Create(value interface{}) error
	Where(query interface{}, args ...interface{}) DBProvider
	First(dest interface{}) error
	Save(value interface{}) error
}

func NewGormPreferenceRepository(db DBProvider) *GormPreferenceRepository {
	return &GormPreferenceRepository{DB: db}
}

func (r *GormPreferenceRepository) GetByUserID(userID uuid.UUID) (*Preference, error) {
	var pref Preference
	err := r.DB.Where("user_id = ?", userID).First(&pref)
	if err != nil {
		return nil, err
	}
	return &pref, nil
}

func (r *GormPreferenceRepository) Set(userID uuid.UUID, inApp, push, email bool) error {
	pref := &Preference{
		UserID: userID,
		InApp:  inApp,
		Push:   push,
		Email:  email,
	}
	return r.DB.Save(pref)
}
