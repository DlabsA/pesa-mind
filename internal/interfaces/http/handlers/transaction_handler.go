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
	// TODO: Get userID from context (after auth)
	userID := uuid.MustParse("00000000-0000-0000-0000-000000000000") // placeholder
	profileID, _ := uuid.Parse(req.ProfileID)
	categoryID, _ := uuid.Parse(req.CategoryID)
	tx, err := h.Service.Create(userID, profileID, categoryID, req.Amount, req.Type, req.Note, req.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.TransactionResponse{
		ID:         tx.ID.String(),
		UserID:     tx.UserID.String(),
		ProfileID:  tx.ProfileID.String(),
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
			UserID:     tx.UserID.String(),
			ProfileID:  tx.ProfileID.String(),
			CategoryID: tx.CategoryID.String(),
			Amount:     tx.Amount,
			Type:       tx.Type,
			Note:       tx.Note,
			Date:       tx.Date.Unix(),
		})
	}
	c.JSON(http.StatusOK, resp)
}
