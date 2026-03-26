package handlers

import (
	"net/http"
	"pesa-mind/internal/domain/automation"
	"pesa-mind/internal/interfaces/http/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AutomationHandler struct {
	Service *automation.Service
}

func NewAutomationHandler(s *automation.Service) *AutomationHandler {
	return &AutomationHandler{Service: s}
}

func (h *AutomationHandler) CreateRule(c *gin.Context) {
	var req dto.CreateSMSAutomationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
	rule, err := h.Service.CreateRule(userID, req.Pattern)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SMSAutomationResponse{
		ID:        rule.ID.String(),
		UserID:    rule.UserID.String(),
		Pattern:   rule.Pattern,
		Enabled:   rule.Enabled,
		CreatedAt: rule.CreatedAt.Unix(),
		UpdatedAt: rule.UpdatedAt.Unix(),
	})
}

func (h *AutomationHandler) ListRules(c *gin.Context) {
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
	rules, err := h.Service.ListRules(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]dto.SMSAutomationResponse, 0, len(rules))
	for _, rule := range rules {
		resp = append(resp, dto.SMSAutomationResponse{
			ID:        rule.ID.String(),
			UserID:    rule.UserID.String(),
			Pattern:   rule.Pattern,
			Enabled:   rule.Enabled,
			CreatedAt: rule.CreatedAt.Unix(),
			UpdatedAt: rule.UpdatedAt.Unix(),
		})
	}
	c.JSON(http.StatusOK, resp)
}
