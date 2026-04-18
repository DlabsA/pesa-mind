package test

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func setupRegisterTestRouter(t *testing.T) (*gin.Engine, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Create tables manually for SQLite compatibility
	db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME,
			email TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL
		)`)
	db.Exec(`
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
		)`)

	userRepo := user.NewGormUserRepository(db)
	userService := user.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router := gin.New()
	router.POST("/users/register", userHandler.Register)

	return router, db
}

func TestRegisterWithCustomUsername(t *testing.T) {
	router, _ := setupRegisterTestRouter(t)

	req := dto.RegisterRequest{
		Email:    "john@example.com",
		Password: "SecurePassword123",
		Username: "johnsmith",
	}

	reqBody, _ := json.Marshal(req)
	request := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewReader(reqBody))
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	var response dto.UserResponse
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(t, "john@example.com", response.Email)
	assert.NotNil(t, response.Profile)
	assert.Equal(t, "johnsmith", response.Profile.Username) // Custom username used
}

func TestRegisterWithDefaultUsername(t *testing.T) {
	router, _ := setupRegisterTestRouter(t)

	req := dto.RegisterRequest{
		Email:    "jane@example.com",
		Password: "SecurePassword123",
		Username: "", // Empty username - should default to email
	}

	reqBody, _ := json.Marshal(req)
	request := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewReader(reqBody))
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	var response dto.UserResponse
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(t, "jane@example.com", response.Email)
	assert.NotNil(t, response.Profile)
	assert.Equal(t, "jane@example.com", response.Profile.Username) // Defaults to email
}

func TestRegisterWithoutUsernameField(t *testing.T) {
	router, _ := setupRegisterTestRouter(t)

	// Request without username field at all
	reqBody := []byte(`{
		"email": "alice@example.com",
		"password": "SecurePassword123"
	}`)
	request := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewReader(reqBody))
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	var response dto.UserResponse
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(t, "alice@example.com", response.Email)
	assert.NotNil(t, response.Profile)
	assert.Equal(t, "alice@example.com", response.Profile.Username) // Defaults to email
}

func TestRegisterInvalidUsernameLength(t *testing.T) {
	router, _ := setupRegisterTestRouter(t)

	req := dto.RegisterRequest{
		Email:    "bob@example.com",
		Password: "SecurePassword123",
		Username: "ab", // Too short (min=3)
	}

	reqBody, _ := json.Marshal(req)
	request := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewReader(reqBody))
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestRegisterUsernameValidation(t *testing.T) {
	tests := []struct {
		name        string
		username    string
		shouldPass  bool
		description string
	}{
		{"Valid short username", "abc", true, "3 chars minimum"},
		{"Valid medium username", "johndoe", true, "Regular username"},
		{"Valid long username", "a" + string(make([]byte, 49)), true, "50 chars max"},
		{"Too short username", "ab", false, "Less than 3 chars"},
		{"Empty username", "", true, "Empty means default to email"},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create fresh router and database for each subtest
			router, _ := setupRegisterTestRouter(t)

			req := dto.RegisterRequest{
				Email:    fmt.Sprintf("test%d@example.com", i),
				Password: "SecurePassword123",
				Username: tc.username,
			}

			reqBody, _ := json.Marshal(req)
			request := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewReader(reqBody))
			request.Header.Set("Content-Type", "application/json")

			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, request)

			if tc.shouldPass {
				assert.Equal(t, http.StatusCreated, recorder.Code, tc.description)
			} else {
				assert.Equal(t, http.StatusBadRequest, recorder.Code, tc.description)
			}
		})
	}
}
