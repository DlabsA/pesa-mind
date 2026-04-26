package handlers

import (
	"net/http"
	"pesa-mind/internal/domain/analytics"
	"pesa-mind/internal/interfaces/http/dto"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AnalyticsHandler struct {
	Service         *analytics.Service
	RealtimeService *analytics.AnalyticsService
}

func NewAnalyticsHandler(s *analytics.Service, realtime *analytics.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{Service: s, RealtimeService: realtime}
}

func (h *AnalyticsHandler) Create(c *gin.Context) {
	var req dto.CreateAnalyticsSnapshotRequest
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
	snap, err := h.Service.Create(userID, req.Type, req.Data, req.Period)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.AnalyticsSnapshotResponse{
		ID:        snap.ID.String(),
		UserID:    snap.UserID.String(),
		Type:      snap.Type,
		Data:      snap.Data,
		Period:    snap.Period,
		CreatedAt: snap.CreatedAt.Unix(),
		UpdatedAt: snap.UpdatedAt.Unix(),
	})
}

func (h *AnalyticsHandler) List(c *gin.Context) {
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
	snaps, err := h.Service.List(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.AnalyticsSnapshotResponse
	for _, snap := range snaps {
		resp = append(resp, dto.AnalyticsSnapshotResponse{
			ID:        snap.ID.String(),
			UserID:    snap.UserID.String(),
			Type:      snap.Type,
			Data:      snap.Data,
			Period:    snap.Period,
			CreatedAt: snap.CreatedAt.Unix(),
			UpdatedAt: snap.UpdatedAt.Unix(),
		})
	}
	c.JSON(http.StatusOK, resp)
}

func (h *AnalyticsHandler) TotalIncome(c *gin.Context) {
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
	amount, err := h.RealtimeService.TotalIncome(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_income": amount})
}

func (h *AnalyticsHandler) TotalExpenses(c *gin.Context) {
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
	amount, err := h.RealtimeService.TotalExpenses(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total_expenses": amount})
}

func (h *AnalyticsHandler) BudgetUtilization(c *gin.Context) {
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
	now := time.Now().UTC()
	util, err := h.RealtimeService.BudgetUtilization(userID, now)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"budget_utilization": util})
}

func (h *AnalyticsHandler) SavingsProgress(c *gin.Context) {
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
	progress, err := h.RealtimeService.SavingsProgress(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"savings_progress": progress})
}
