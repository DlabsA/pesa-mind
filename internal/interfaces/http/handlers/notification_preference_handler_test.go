package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pesa-mind/internal/domain/notification"
)

type MockPreferenceService struct {
	mock.Mock
}

func (m *MockPreferenceService) Get(userID uuid.UUID) (*notification.Preference, error) {
	args := m.Called(userID)
	if pref, ok := args.Get(0).(*notification.Preference); ok {
		return pref, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockPreferenceService) Set(userID uuid.UUID, inApp, push, email bool) error {
	args := m.Called(userID, inApp, push, email)
	return args.Error(0)
}

func TestNotificationPreferenceHandler_Get_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockPreferenceService)
	h := NewNotificationPreferenceHandler(mockService)
	r := gin.New()
	r.GET("/pref", func(c *gin.Context) {
		c.Set("user_id", uuid.New().String())
		h.Get(c)
	})

	mockService.On("Get", mock.Anything).Return(&notification.Preference{
		InApp: true, Push: false, Email: true,
	}, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pref", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "in_app")
}

func TestNotificationPreferenceHandler_Get_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewNotificationPreferenceHandler(new(MockPreferenceService))
	r := gin.New()
	r.GET("/pref", h.Get)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pref", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestNotificationPreferenceHandler_Set_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockPreferenceService)
	h := NewNotificationPreferenceHandler(mockService)
	r := gin.New()
	r.POST("/pref", func(c *gin.Context) {
		c.Set("user_id", uuid.New().String())
		h.Set(c)
	})

	mockService.On("Set", mock.Anything, true, false, true).Return(nil)

	body := `{"in_app":true,"push":false,"email":true}`
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/pref", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "updated")
}

func TestNotificationPreferenceHandler_Set_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewNotificationPreferenceHandler(new(MockPreferenceService))
	r := gin.New()
	r.POST("/pref", func(c *gin.Context) {
		c.Set("user_id", uuid.New().String())
		h.Set(c)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/pref", strings.NewReader("invalid-json"))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestNotificationPreferenceHandler_Set_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	h := NewNotificationPreferenceHandler(new(MockPreferenceService))
	r := gin.New()
	r.POST("/pref", h.Set)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/pref", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestNotificationPreferenceHandler_Get_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockPreferenceService)
	h := NewNotificationPreferenceHandler(mockService)
	r := gin.New()
	r.GET("/pref", func(c *gin.Context) {
		c.Set("user_id", uuid.New().String())
		h.Get(c)
	})

	mockService.On("Get", mock.Anything).Return(nil, errors.New("db error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/pref", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
