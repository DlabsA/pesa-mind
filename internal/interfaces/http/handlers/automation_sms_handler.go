package handlers

import (
	"net/http"
	"pesa-mind/internal/domain/transaction"

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
	UserID           string  `json:"user_id" binding:"required,uuid4"`
	Amount           float64 `json:"amount" binding:"required"`
	ChannelDetailsID string  `json:"channel_details_id" binding:"required,uuid4"`
	Type             string  `json:"type" binding:"required"`
	Note             string  `json:"note"`
	Description      string  `json:"description" binding:"required"`
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
	ChannelDetailsID, _ := uuid.Parse(payload.ChannelDetailsID)
	channel, err := h.TransactionService.Category.GetByID(ChannelDetailsID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get channel details"})
		return
	}
	if channel == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "channel details not found"})
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
	user, _, err := h.TransactionService.User.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get profile"})
		return
	}
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "profile not found"})
		return
	}
	ChannelDetailsID, _ = uuid.Parse(payload.ChannelDetailsID)
	channel, err = h.TransactionService.Category.GetByID(ChannelDetailsID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get channel details"})
		return
	}
	if channel == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "channel details not found"})
		return
	}

	tx, err := h.TransactionService.CreateTransactionFromAutomation(user, payload.Amount, channel, payload.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"transaction_id": tx.ID})
}
