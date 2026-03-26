package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pesa-mind/internal/domain/user"
	"pesa-mind/internal/interfaces/http/dto"
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
	// Password hashing and validation will be handled in Auth
	user, err := h.Service.Register(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.UserResponse{ID: user.ID.String(), Email: user.Email})
}
