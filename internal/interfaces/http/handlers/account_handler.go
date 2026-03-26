package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"pesa-mind/internal/domain/account"
	"pesa-mind/internal/interfaces/http/dto"
)

type AccountHandler struct {
	Service *account.Service
}

func NewAccountHandler(s *account.Service) *AccountHandler {
	return &AccountHandler{Service: s}
}

func (h *AccountHandler) Create(c *gin.Context) {
	var req dto.CreateAccountRequest
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
	account, err := h.Service.Create(userID, req.Name, req.Type, req.Currency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.AccountResponse{
		ID:       account.ID.String(),
		UserID:   account.UserID.String(),
		Name:     account.Name,
		Type:     account.Type,
		Currency: account.Currency,
		Balance:  account.Balance,
	})
}

func (h *AccountHandler) List(c *gin.Context) {
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
	accounts, err := h.Service.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.AccountResponse
	for _, a := range accounts {
		resp = append(resp, dto.AccountResponse{
			ID:       a.ID.String(),
			UserID:   a.UserID.String(),
			Name:     a.Name,
			Type:     a.Type,
			Currency: a.Currency,
			Balance:  a.Balance,
		})
	}
	c.JSON(http.StatusOK, resp)
}
