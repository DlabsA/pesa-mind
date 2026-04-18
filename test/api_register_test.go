package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"pesa-mind/internal/domain/user"
	"pesa-mind/internal/interfaces/http/dto"
	"pesa-mind/internal/interfaces/http/handlers"
)

func setupTestRouter(t *testing.T) (*gin.Engine, *gorm.DB) {
	// SQLite doesn't support uuid_generate_v4, use postgres driver for testing or skip UUID features
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err, "failed to create test database")

	// Manually create tables for SQLite (bypassing uuid issues)
	err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME,
			email TEXT NOT NULL,
			password_hash TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS profiles (
			id TEXT PRIMARY KEY,
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME,
			user_id TEXT NOT NULL UNIQUE,
			username TEXT NOT NULL UNIQUE,
			type TEXT DEFAULT 'Free',
			balance REAL DEFAULT 0.0,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);
	`).Error
	assert.NoError(t, err, "failed to create test tables")

	// Initialize services and handlers
	userRepo := user.NewGormUserRepository(db)
	userService := user.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Setup router
	router := gin.New()
	router.POST("/users/register", userHandler.Register)

	return router, db
}

func TestAPIRegisterWithProfile(t *testing.T) {
	router, db := setupTestRouter(t)

	// Create registration request
	req := dto.RegisterRequest{
		Email:    "api@example.com",
		Password: "securepassword123",
	}

	reqBody, _ := json.Marshal(req)
	request := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewReader(reqBody))
	request.Header.Set("Content-Type", "application/json")

	// Create response recorder
	recorder := httptest.NewRecorder()

	// Perform request
	router.ServeHTTP(recorder, request)

	// Assertions
	assert.Equal(t, http.StatusCreated, recorder.Code, "should return 201 Created")

	// Parse response
	var response dto.UserResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err, "response should be valid JSON")
	assert.Equal(t, "api@example.com", response.Email)
	assert.NotEmpty(t, response.ID, "user ID should be present")
	assert.NotNil(t, response.Profile, "profile should be included in response")
	assert.Equal(t, "api@example.com", response.Profile.Username)
	assert.Equal(t, "Free", response.Profile.Type)
	assert.Equal(t, 0.0, response.Profile.Balance)

	// Verify in database
	var profileCount int64
	db.Model(&user.Profile{}).Count(&profileCount)
	assert.Equal(t, int64(1), profileCount, "profile should be created in database")
}

func TestAPIRegisterInvalidEmail(t *testing.T) {
	router, _ := setupTestRouter(t)

	req := dto.RegisterRequest{
		Email:    "invalid-email",
		Password: "password123",
	}

	reqBody, _ := json.Marshal(req)
	request := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewReader(reqBody))
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code, "should return 400 for invalid email")
}

func TestAPIRegisterShortPassword(t *testing.T) {
	router, _ := setupTestRouter(t)

	req := dto.RegisterRequest{
		Email:    "valid@example.com",
		Password: "short",
	}

	reqBody, _ := json.Marshal(req)
	request := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewReader(reqBody))
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code, "should return 400 for short password")
}

func TestAPIRegisterMultipleUsers(t *testing.T) {
	router, db := setupTestRouter(t)

	// Register first user
	req1 := dto.RegisterRequest{
		Email:    "user1@example.com",
		Password: "password123",
	}
	reqBody1, _ := json.Marshal(req1)
	request1 := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewReader(reqBody1))
	request1.Header.Set("Content-Type", "application/json")
	recorder1 := httptest.NewRecorder()
	router.ServeHTTP(recorder1, request1)
	assert.Equal(t, http.StatusCreated, recorder1.Code)

	// Register second user
	req2 := dto.RegisterRequest{
		Email:    "user2@example.com",
		Password: "password456",
	}
	reqBody2, _ := json.Marshal(req2)
	request2 := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewReader(reqBody2))
	request2.Header.Set("Content-Type", "application/json")
	recorder2 := httptest.NewRecorder()
	router.ServeHTTP(recorder2, request2)
	assert.Equal(t, http.StatusCreated, recorder2.Code)

	// Verify both users and profiles exist
	var userCount int64
	db.Model(&user.User{}).Count(&userCount)
	assert.Equal(t, int64(2), userCount, "should have 2 users")

	var profileCount int64
	db.Model(&user.Profile{}).Count(&profileCount)
	assert.Equal(t, int64(2), profileCount, "should have 2 profiles")
}
