package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// TestUserProfileModel verifies the User and Profile models work correctly
func TestUserProfileModel(t *testing.T) {
	// Create a user
	user := &User{
		Email:        "test@example.com",
		PasswordHash: "hashed_password",
	}
	user.ID = uuid.New()

	assert.Equal(t, "test@example.com", user.Email)
	assert.NotEmpty(t, user.ID)

	// Create a profile
	profile := &Profile{
		UserID:   user.ID,
		Username: "testuser",
		Type:     "Free",
		Balance:  0.0,
	}
	profile.ID = uuid.New()

	assert.Equal(t, "testuser", profile.Username)
	assert.Equal(t, user.ID, profile.UserID)
	assert.Equal(t, "Free", profile.Type)
	assert.Equal(t, 0.0, profile.Balance)
}

// TestUserProfileRelationship verifies the relationship between User and Profile
func TestUserProfileRelationship(t *testing.T) {
	userID := uuid.New()

	user := &User{
		Email:        "related@example.com",
		PasswordHash: "pass123",
	}
	user.ID = userID

	profile := &Profile{
		UserID:   userID,
		Username: "related@example.com",
		Type:     "Premium",
		Balance:  100.0,
	}
	profile.ID = uuid.New()

	// Set user's profile pointer
	user.Profile = profile

	assert.NotNil(t, user.Profile)
	assert.Equal(t, user.Profile.UserID, user.ID)
}

// TestUserProfileDefaults verifies default values
func TestUserProfileDefaults(t *testing.T) {
	profile := &Profile{
		UserID:   uuid.New(),
		Username: "defaultuser",
		Type:     "Free",
		Balance:  0.0,
	}
	profile.ID = uuid.New()

	assert.Equal(t, "Free", profile.Type)
	assert.Equal(t, 0.0, profile.Balance)
}

// TestMultipleProfiles verifies multiple profiles can exist
func TestMultipleProfiles(t *testing.T) {
	profiles := make([]Profile, 3)

	for i := 0; i < 3; i++ {
		profiles[i] = Profile{
			UserID:   uuid.New(),
			Username: "user" + string(rune(i)),
			Type:     "Free",
			Balance:  0.0,
		}
		profiles[i].ID = uuid.New()
	}

	assert.Equal(t, 3, len(profiles))
	for i := 0; i < 3; i++ {
		assert.NotEqual(t, profiles[i].ID, profiles[(i+1)%3].ID)
		assert.NotEqual(t, profiles[i].UserID, profiles[(i+1)%3].UserID)
	}
}

// TestUserProfileLifecycle verifies the complete lifecycle
func TestUserProfileLifecycle(t *testing.T) {
	// 1. Create user
	user := &User{
		Email:        "lifecycle@example.com",
		PasswordHash: "secure_hash",
	}
	user.ID = uuid.New()

	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "lifecycle@example.com", user.Email)

	// 2. Create profile
	profile := &Profile{
		UserID:   user.ID,
		Username: "lifecycle@example.com",
		Type:     "Free",
		Balance:  0.0,
	}
	profile.ID = uuid.New()

	assert.Equal(t, user.ID, profile.UserID)

	// 3. Link profile to user
	user.Profile = profile
	profile.User = user

	assert.Equal(t, user.Profile.ID, profile.ID)
	assert.Equal(t, profile.User.Email, user.Email)

	// 4. Update balance
	profile.Balance = 50.0
	assert.Equal(t, 50.0, user.Profile.Balance)

	// 5. Update type
	profile.Type = "Premium"
	assert.Equal(t, "Premium", user.Profile.Type)
}
