package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"pesa-mind/internal/domain/gamification"
	"pesa-mind/internal/interfaces/http/dto"
)

type GamificationHandler struct {
	Service *gamification.Service
}

func NewGamificationHandler(s *gamification.Service) *GamificationHandler {
	return &GamificationHandler{Service: s}
}

func (h *GamificationHandler) ListBadges(c *gin.Context) {
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
	badges, err := h.Service.GetUserBadges(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]dto.BadgeResponse, 0, len(badges))
	for _, b := range badges {
		resp = append(resp, dto.BadgeResponse{
			ID:       b.BadgeID.String(),
			EarnedAt: b.EarnedAt.Unix(),
		})
	}
	c.JSON(http.StatusOK, resp)
}

func (h *GamificationHandler) ListStreaks(c *gin.Context) {
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
	streaks, err := h.Service.GetUserStreaks(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]dto.StreakResponse, 0, len(streaks))
	for _, s := range streaks {
		resp = append(resp, dto.StreakResponse{
			Type:   s.Type,
			Count:  s.Count,
			Active: true, // Could add logic for active/inactive
		})
	}
	c.JSON(http.StatusOK, resp)
}

func (h *GamificationHandler) ListAchievements(c *gin.Context) {
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
	achievements, err := h.Service.GetUserAchievements(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]dto.AchievementResponse, 0, len(achievements))
	for _, a := range achievements {
		resp = append(resp, dto.AchievementResponse{
			ID:       a.AchievementID.String(),
			EarnedAt: a.EarnedAt.Unix(),
		})
	}
	c.JSON(http.StatusOK, resp)
}

func (h *GamificationHandler) GetLeaderboard(c *gin.Context) {
	period := c.DefaultQuery("period", "weekly")
	limit := 10
	entries, err := h.Service.GetLeaderboard(period, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]dto.LeaderboardEntryResponse, 0, len(entries))
	for _, e := range entries {
		resp = append(resp, dto.LeaderboardEntryResponse{
			UserID: e.UserID.String(),
			Score:  e.Score,
			Rank:   e.Rank,
		})
	}
	c.JSON(http.StatusOK, resp)
}

func (h *GamificationHandler) ListRewards(c *gin.Context) {
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
	rewards, err := h.Service.GetUserRewards(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]dto.RewardResponse, 0, len(rewards))
	for _, r := range rewards {
		resp = append(resp, dto.RewardResponse{
			ID:        r.RewardID.String(),
			ClaimedAt: r.ClaimedAt.Unix(),
		})
	}
	c.JSON(http.StatusOK, resp)
}

func (h *GamificationHandler) ClaimReward(c *gin.Context) {
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
	rewardIDStr := c.Param("reward_id")
	rewardID, err := uuid.Parse(rewardIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid reward id"})
		return
	}
	if err := h.Service.ClaimReward(userID, rewardID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "reward claimed"})
}
