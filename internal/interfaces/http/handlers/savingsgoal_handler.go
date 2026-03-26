package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"pesa-mind/internal/domain/savingsgoal"
	"pesa-mind/internal/interfaces/http/dto"
	"strconv"
)

type SavingsGoalHandler struct {
	Service *savingsgoal.Service
}

func NewSavingsGoalHandler(s *savingsgoal.Service) *SavingsGoalHandler {
	return &SavingsGoalHandler{Service: s}
}

func (h *SavingsGoalHandler) Create(c *gin.Context) {
	var req dto.CreateSavingsGoalRequest
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
	goal, err := h.Service.Create(userID, req.Title, req.Target, req.Deadline, req.AutoSave)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var deadlineUnix *int64
	if goal.Deadline != nil {
		d := goal.Deadline.Unix()
		deadlineUnix = &d
	}
	c.JSON(http.StatusCreated, dto.SavingsGoalResponse{
		ID:       goal.ID.String(),
		UserID:   goal.UserID.String(),
		Title:    goal.Title,
		Target:   goal.Target,
		Current:  goal.Current,
		Deadline: deadlineUnix,
		AutoSave: goal.AutoSave,
	})
}

func (h *SavingsGoalHandler) List(c *gin.Context) {
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
	goals, err := h.Service.List(userID, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var resp []dto.SavingsGoalResponse
	for _, goal := range goals {
		var deadlineUnix *int64
		if goal.Deadline != nil {
			d := goal.Deadline.Unix()
			deadlineUnix = &d
		}
		resp = append(resp, dto.SavingsGoalResponse{
			ID:       goal.ID.String(),
			UserID:   goal.UserID.String(),
			Title:    goal.Title,
			Target:   goal.Target,
			Current:  goal.Current,
			Deadline: deadlineUnix,
			AutoSave: goal.AutoSave,
		})
	}
	c.JSON(http.StatusOK, resp)
}
