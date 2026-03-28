package http

import (
	"github.com/gin-gonic/gin"
	"pesa-mind/internal/interfaces/http/handlers"
)

func RegisterRoutes(r *gin.Engine, deps *handlers.Dependencies) {
	api := r.Group("/api/v1")

	// ...existing routes...

	budget := handlers.NewBudgetHandler(deps.BudgetService)
	api.POST("/budgets", budget.Create)
	api.GET("/budgets", budget.List)

	analytics := handlers.NewAnalyticsHandler(deps.AnalyticsServiceCRUD, deps.AnalyticsService)
	api.GET("/analytics/income", analytics.TotalIncome)
	api.GET("/analytics/expenses", analytics.TotalExpenses)
	api.GET("/analytics/budget-utilization", analytics.BudgetUtilization)
	api.GET("/analytics/savings-progress", analytics.SavingsProgress)

	// Gamification endpoints
	gamification := handlers.NewGamificationHandler(deps.GamificationService)
	api.GET("/gamification/badges", gamification.ListBadges)
	api.GET("/gamification/streaks", gamification.ListStreaks)
	api.GET("/gamification/achievements", gamification.ListAchievements)
	api.GET("/gamification/leaderboard", gamification.GetLeaderboard)
	api.GET("/gamification/rewards", gamification.ListRewards)
	api.POST("/gamification/rewards/:reward_id/claim", gamification.ClaimReward)
}
