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
	_, profile, err := h.Service.Profile.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get profile"})
		return
	}
	if profile == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "profile not found"})
		return
	}

	categoryID, _ := uuid.Parse(req.CategoryID)
	tx, err := h.Service.Create(profile, categoryID, req.Amount, req.Type, req.Note, req.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.TransactionResponse{
		ID:         tx.ID.String(),
		ProfileID:  tx.Profile.ID.String(),
		CategoryID: tx.CategoryID.String(),
		Amount:     tx.Amount,
		Type:       tx.Type,
		Note:       tx.Note,
		Date:       tx.Date.Unix(),
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
			ID:         tx.ID.String(),
			ProfileID:  tx.Profile.ID.String(),
			CategoryID: tx.CategoryID.String(),
			Amount:     tx.Amount,
			Type:       tx.Type,
			Note:       tx.Note,
			Date:       tx.Date.Unix(),
		})
	}
	c.JSON(http.StatusOK, resp)
}
