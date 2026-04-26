package handlers

import (
	"net/http"
	"pesa-mind/internal/domain/transaction"
	"pesa-mind/internal/interfaces/http/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionHandler struct {
	Service *transaction.Service
}

func NewTransactionHandler(s *transaction.Service) *TransactionHandler {
	return &TransactionHandler{Service: s}
}

func (h *TransactionHandler) Create(c *gin.Context) {
	var req dto.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get userID from context
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

	// Get the user's profile
	user, _, err := h.Service.User.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get profile"})
		return
	}
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "profile not found"})
		return
	}

	ChannelDetailsID, _ := uuid.Parse(req.ChannelDetailsID)
	channel, err := h.Service.Category.GetByID(ChannelDetailsID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get channel details"})
		return
	}
	if channel == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "channel details not found"})
		return
	}
	tx, err := h.Service.Create(user, channel, req.Amount, req.Type, req.Note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.TransactionResponse{
		ID:                 tx.ID.String(),
		Username:           user.Profile.Username,
		ChannelDetailsName: channel.Name,
		Amount:             tx.Amount,
		Type:               tx.Type,
		Note:               tx.Note,
	})
}

func (h *TransactionHandler) List(c *gin.Context) {
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
	txs, err := h.Service.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.TransactionResponse
	for _, tx := range txs {
		resp = append(resp, dto.TransactionResponse{
			ID:                 tx.ID.String(),
			Username:           tx.User.Profile.Username,
			ChannelDetailsName: tx.ChannelDetails.Name,
			Amount:             tx.Amount,
			Type:               tx.Type,
			Note:               tx.Note,
		})
	}
	c.JSON(http.StatusOK, resp)
}
