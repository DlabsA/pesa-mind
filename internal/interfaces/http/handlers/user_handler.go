package handlers

import (
	"errors"
	"net/http"
	"pesa-mind/internal/domain/user"
	"pesa-mind/internal/interfaces/http/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	Service *user.Service
}

func getUserIDFromContextOrParam(c *gin.Context) (string, bool) {
	// First try param (keeps compatibility if a :id param is used)
	if id := c.Param("id"); id != "" {
		return id, true
	}
	// Then try middleware-set user_id
	if v, ok := c.Get("user_id"); ok {
		if s, ok := v.(string); ok && s != "" {
			return s, true
		}
	}
	return "", false
}

func NewUserHandler(s *user.Service) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Hash the password before saving
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	// Pass username (will default to email in service if empty)
	usr, err := h.Service.Register(req.Email, string(hashed), req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := dto.UserResponse{
		ID:    usr.ID.String(),
		Email: usr.Email,
	}

	if usr.Profile != nil {
		resp.Profile = &dto.ProfileData{
			ID:       usr.Profile.ID.String(),
			UserID:   usr.Profile.UserID.String(),
			Username: usr.Profile.Username,
			Type:     usr.Profile.Type,
			Balance:  usr.Profile.Balance,
		}
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *UserHandler) Get(c *gin.Context) {
	idStr, ok := getUserIDFromContextOrParam(c)
	if !ok || idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	uid, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	u, profile, err := h.Service.GetByID(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if u == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	selectedProfile := profile
	if selectedProfile == nil {
		selectedProfile = u.Profile
	}

	resp := dto.UserResponse{
		ID:      u.ID.String(),
		Email:   u.Email,
		Profile: dto.ToProfileDTO(selectedProfile),
	}
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) Update(c *gin.Context) {
	idStr, ok := getUserIDFromContextOrParam(c)
	if !ok || idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	uid, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service.Update with the id + fields to change (service will load and persist)
	updatedUser, updatedProfile, err := h.Service.Update(uid, req.Email, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if updatedUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// Use DTO mapper to build response
	resp := dto.ToUserResponse(updatedUser, updatedProfile)
	c.JSON(http.StatusOK, resp)
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	idStr, ok := getUserIDFromContextOrParam(c)
	if !ok || idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	uid, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to verify current password and set new one
	updatedUser, err := h.Service.ChangePassword(uid, req.CurrentPassword, req.NewPassword)
	if err != nil {
		// check sentinel error for invalid current password
		if errors.Is(err, user.ErrInvalidCurrentPassword) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "current password is incorrect"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	if updatedUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// don't return user with passwordhash — return minimal response (ok)
	c.JSON(http.StatusOK, gin.H{"message": "password changed successfully"})
}
