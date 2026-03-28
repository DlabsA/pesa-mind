package handlers

import (
	"net/http"
	"pesa-mind/internal/domain/notification"
	"pesa-mind/internal/interfaces/http/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PreferenceService interface {
	Get(userID uuid.UUID) (*notification.Preference, error)
	Set(userID uuid.UUID, inApp, push, email bool) error
}

type NotificationPreferenceHandler struct {
	Service PreferenceService
}

func NewNotificationPreferenceHandler(s PreferenceService) *NotificationPreferenceHandler {
	return &NotificationPreferenceHandler{Service: s}
}

func (h *NotificationPreferenceHandler) Get(c *gin.Context) {
	userIDStr, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	pref, err := h.Service.Get(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.NotificationPreferenceResponse{
		InApp: pref.InApp,
		Push:  pref.Push,
		Email: pref.Email,
	})
}

func (h *NotificationPreferenceHandler) Set(c *gin.Context) {
	userIDStr, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, err := uuid.Parse(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	var req dto.SetNotificationPreferenceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Set(userID, req.InApp, req.Push, req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}
