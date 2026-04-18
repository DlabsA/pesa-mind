package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestRepositoryCreateUserAndProfile(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	db.Exec(`CREATE TABLE users (
		id TEXT PRIMARY KEY,
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL
	)`)
	db.Exec(`CREATE TABLE profiles (
		id TEXT PRIMARY KEY,
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME,
		user_id TEXT NOT NULL UNIQUE,
		username TEXT NOT NULL UNIQUE,
		type TEXT NOT NULL,
		balance REAL DEFAULT 0.0
	)`)

	repo := NewGormUserRepository(db)
	userID := uuid.New()

	userProfile := UserProfile{
		user: &User{
			Email:        "test@example.com",
			PasswordHash: "hashedpass",
		},
		username: "testuser",
	}
	userProfile.user.ID = userID

	err := repo.Create(userProfile)
	assert.NoError(t, err, "Create should succeed")

	// Verify user exists
	var user User
	db.First(&user, "email = ?", "test@example.com")
	assert.Equal(t, "test@example.com", user.Email)

	// Verify profile exists
	var profile Profile
	db.First(&profile, "username = ?", "testuser")
	assert.Equal(t, "testuser", profile.Username)
	assert.Equal(t, "Free", profile.Type)
}

func TestRepositoryFindByEmail(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	db.Exec(`CREATE TABLE users (
		id TEXT PRIMARY KEY,
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL
	)`)
	db.Exec(`CREATE TABLE profiles (
		id TEXT PRIMARY KEY,
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME,
		user_id TEXT NOT NULL UNIQUE,
		username TEXT NOT NULL UNIQUE,
		type TEXT,
		balance REAL
	)`)

	userID := uuid.New()
	user := User{Email: "find@test.com", PasswordHash: "hash"}
	user.ID = userID
	db.Create(&user)

	profile := Profile{UserID: userID, Username: "finduser", Type: "Premium", Balance: 100.0}
	profile.ID = uuid.New()
	db.Create(&profile)

	repo := NewGormUserRepository(db)
	foundUser, foundProfile, err := repo.FindByEmail("find@test.com")

	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.NotNil(t, foundProfile)
	assert.Equal(t, "find@test.com", foundUser.Email)
	assert.Equal(t, "finduser", foundProfile.Username)
	assert.Equal(t, 100.0, foundProfile.Balance)
}

func TestServiceRegister(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	db.Exec(`CREATE TABLE users (
		id TEXT PRIMARY KEY,
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL
	)`)
	db.Exec(`CREATE TABLE profiles (
		id TEXT PRIMARY KEY,
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME,
		user_id TEXT NOT NULL UNIQUE,
		username TEXT NOT NULL UNIQUE,
		type TEXT,
		balance REAL
	)`)

	repo := NewGormUserRepository(db)
	service := NewService(repo)

	user, err := service.Register("register@test.com", "hashed123", "")
	assert.NoError(t, err, "Register should succeed")
	assert.NotNil(t, user)
	assert.Equal(t, "register@test.com", user.Email)

	// Verify profile was created
	var profile Profile
	db.First(&profile, "user_id = ?", user.ID)
	assert.Equal(t, "register@test.com", profile.Username)
	assert.Equal(t, "Free", profile.Type)
	assert.Equal(t, 0.0, profile.Balance)
}

func TestServiceGetByEmailWithProfile(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	db.Exec(`CREATE TABLE users (
		id TEXT PRIMARY KEY,
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL
	)`)
	db.Exec(`CREATE TABLE profiles (
		id TEXT PRIMARY KEY,
		created_at DATETIME,
		updated_at DATETIME,
		deleted_at DATETIME,
		user_id TEXT NOT NULL UNIQUE,
		username TEXT NOT NULL UNIQUE,
		type TEXT,
		balance REAL
	)`)

	repo := NewGormUserRepository(db)
	service := NewService(repo)

	// Register user
	service.Register("get@test.com", "hash456", "")

	// Get by email
	user, profile, err := service.GetByEmail("get@test.com")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, profile)
	assert.Equal(t, "get@test.com", user.Email)
	assert.Equal(t, "get@test.com", profile.Username)
	assert.Equal(t, "Free", profile.Type)
}
