package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"pesa-mind/internal/domain/budget"
	"pesa-mind/internal/interfaces/http/dto"
	"strconv"
)

type BudgetHandler struct {
	Service *budget.Service
}

func NewBudgetHandler(s *budget.Service) *BudgetHandler {
	return &BudgetHandler{Service: s}
}

func (h *BudgetHandler) Create(c *gin.Context) {
	var req dto.CreateBudgetRequest
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
	b, err := h.Service.Create(userID, req.Name, req.Amount, req.Period, req.StartDate, req.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.BudgetResponse{
		ID:        b.ID.String(),
		UserID:    b.UserID.String(),
		Name:      b.Name,
		Amount:    b.Amount,
		Period:    b.Period,
		StartDate: b.StartDate.Unix(),
		EndDate:   b.EndDate.Unix(),
	})
}

func (h *BudgetHandler) List(c *gin.Context) {
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
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	budgets, err := h.Service.List(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.BudgetResponse
	for _, b := range budgets {
		resp = append(resp, dto.BudgetResponse{
			ID:        b.ID.String(),
			UserID:    b.UserID.String(),
			Name:      b.Name,
			Amount:    b.Amount,
			Period:    b.Period,
			StartDate: b.StartDate.Unix(),
			EndDate:   b.EndDate.Unix(),
		})
	}
	c.JSON(http.StatusOK, resp)
}
