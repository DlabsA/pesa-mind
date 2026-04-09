package handlers

import (
	"net/http"
	"pesa-mind/internal/domain/user"
	"pesa-mind/internal/interfaces/http/dto"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	Service *user.Service
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
