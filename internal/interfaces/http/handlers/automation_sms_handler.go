package handlers

import (
	"net/http"
	"pesa-mind/internal/domain/transaction"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AutomationSMSHandler struct {
	TransactionService *transaction.Service
}

func NewAutomationSMSHandler(ts *transaction.Service) *AutomationSMSHandler {
	return &AutomationSMSHandler{TransactionService: ts}
}

// SMSPayload is the expected payload from the mobile app for an SMS transaction
// Example: {"user_id": "...", "amount": 1000, "category_id": "...", "description": "Airtime purchase", "occurred_at": "2026-03-25T12:34:56Z"}
type SMSPayload struct {
	UserID      string  `json:"user_id" binding:"required,uuid4"`
	Amount      float64 `json:"amount" binding:"required"`
	CategoryID  string  `json:"category_id" binding:"required,uuid4"`
	Description string  `json:"description" binding:"required"`
	OccurredAt  string  `json:"occurred_at" binding:"required,datetime"`
}

// POST /api/v1/automation/sms/transaction
func (h *AutomationSMSHandler) CreateTransactionFromSMS(c *gin.Context) {
	var payload SMSPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := uuid.Parse(payload.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}
	categoryID, err := uuid.Parse(payload.CategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category_id"})
		return
	}
	occurredAt, err := time.Parse(time.RFC3339, payload.OccurredAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid occurred_at"})
		return
	}
	// Get userID from context
	userIDStr, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, err = uuid.Parse(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	// Get the user's profile
	_, profile, err := h.TransactionService.Profile.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get profile"})
		return
	}
	if profile == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "profile not found"})
		return
	}
	tx, err := h.TransactionService.CreateTransactionFromAutomation(profile, payload.Amount, categoryID, payload.Description, occurredAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"transaction_id": tx.ID})
}
