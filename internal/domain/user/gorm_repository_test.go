package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB creates an in-memory SQLite database for testing
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err, "failed to open test database")

	// Create tables manually for SQLite compatibility
	db.Exec(`
		CREATE TABLE users (
			id TEXT PRIMARY KEY,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			deleted_at DATETIME,
			email TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL
		)
	`)

	db.Exec(`
		CREATE TABLE profiles (
			id TEXT PRIMARY KEY,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			deleted_at DATETIME,
			user_id TEXT NOT NULL UNIQUE,
			username TEXT NOT NULL UNIQUE,
			type TEXT NOT NULL,
			balance REAL DEFAULT 0.0,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`)

	return db
}

func TestGormUserRepositoryCreate(t *testing.T) {
	db := setupTestDB(t)
	repo := NewGormUserRepository(db)

	// Create a user with profile
	userProfile := UserProfile{
		user: &User{
			Email:        "test@example.com",
			PasswordHash: "hashedpassword",
		},
		username: "testuser",
	}

	err := repo.Create(userProfile)
	assert.NoError(t, err, "failed to create user with profile")

	// Verify user was created
	var user User
	result := db.First(&user, "email = ?", "test@example.com")
	assert.NoError(t, result.Error, "user should exist in database")
	assert.Equal(t, "test@example.com", user.Email)

	// Verify profile was created
	var profile Profile
	result = db.First(&profile, "user_id = ?", user.ID)
	assert.NoError(t, result.Error, "profile should exist in database")
	assert.Equal(t, "testuser", profile.Username)
	assert.Equal(t, "Free", profile.Type)
	assert.Equal(t, 0.0, profile.Balance)
}

func TestGormUserRepositoryFindByID(t *testing.T) {
	db := setupTestDB(t)
	repo := NewGormUserRepository(db)

	// Create user with profile
	userID := uuid.New()
	user := &User{
		Email:        "find@example.com",
		PasswordHash: "hash123",
	}
	user.ID = userID

	db.Create(user)

	profile := &Profile{
		UserID:   userID,
		Username: "finduser",
		Type:     "Premium",
		Balance:  150.0,
	}
	db.Create(profile)

	// Find by ID
	foundUser, foundProfile, err := repo.FindByID(userID)
	assert.NoError(t, err, "should find user by ID")
	assert.NotNil(t, foundUser, "user should not be nil")
	assert.NotNil(t, foundProfile, "profile should not be nil")
	assert.Equal(t, "find@example.com", foundUser.Email)
	assert.Equal(t, "finduser", foundProfile.Username)
	assert.Equal(t, "Premium", foundProfile.Type)
}

func TestGormUserRepositoryFindByEmail(t *testing.T) {
	db := setupTestDB(t)
	repo := NewGormUserRepository(db)

	// Create user with profile
	userID := uuid.New()
	user := &User{
		Email:        "email@example.com",
		PasswordHash: "hash456",
	}
	user.ID = userID

	db.Create(user)

	profile := &Profile{
		UserID:   userID,
		Username: "emailuser",
		Type:     "Enterprise",
		Balance:  500.0,
	}
	db.Create(profile)

	// Find by email
	foundUser, foundProfile, err := repo.FindByEmail("email@example.com")
	assert.NoError(t, err, "should find user by email")
	assert.NotNil(t, foundUser, "user should not be nil")
	assert.NotNil(t, foundProfile, "profile should not be nil")
	assert.Equal(t, "emailuser", foundProfile.Username)
	assert.Equal(t, "Enterprise", foundProfile.Type)
	assert.Equal(t, 500.0, foundProfile.Balance)
}

func TestGormUserRepositoryUpdate(t *testing.T) {
	db := setupTestDB(t)
	repo := NewGormUserRepository(db)

	// Create user
	user := &User{
		Email:        "update@example.com",
		PasswordHash: "oldhash",
	}
	user.ID = uuid.New()
	db.Create(user)

	// Update user
	user.PasswordHash = "newhash"
	err := repo.Update(user)
	assert.NoError(t, err, "should update user")

	// Verify update
	var updated User
	db.First(&updated, "id = ?", user.ID)
	assert.Equal(t, "newhash", updated.PasswordHash)
}

func TestGormUserRepositoryDelete(t *testing.T) {
	db := setupTestDB(t)
	repo := NewGormUserRepository(db)

	// Create user
	user := &User{
		Email:        "delete@example.com",
		PasswordHash: "hash",
	}
	user.ID = uuid.New()
	db.Create(user)

	// Delete user
	err := repo.Delete(user.ID)
	assert.NoError(t, err, "should delete user")

	// Verify deletion (soft delete by default)
	var deleted User
	result := db.First(&deleted, "id = ?", user.ID)
	assert.Error(t, result.Error, "user should be soft deleted")
}
