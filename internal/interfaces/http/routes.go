package http

import (
	"pesa-mind/internal/interfaces/http/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, deps *handlers.Dependencies) {
	api := r.Group("/api/v1")

	// ...existing routes...

	budget := handlers.NewBudgetHandler(deps.BudgetService)
	// Budget endpoints
	api.POST("/budgets", budget.Create)
	api.GET("/budgets", budget.List)
	api.GET("/budgets/:id", budget.GetByID)
	api.PUT("/budgets/:id", budget.Update)
	api.DELETE("/budgets/:id", budget.Delete)

	// MonthlyBudget endpoints
	api.POST("/monthly-budgets", budget.CreateMonthlyBudget)
	api.GET("/monthly-budgets/:id", budget.GetMonthlyBudgetByID)
	api.GET("/monthly-budgets", budget.ListMonthlyBudgets)
	api.GET("/yearly-budgets/:yearly_budget_id/monthly", budget.GetMonthlyBudgetsByYearlyBudgetID)
	api.GET("/monthly-budgets/by-month", budget.GetMonthlyBudgetByUserIDAndMonthYear)
	api.PUT("/monthly-budgets/:id", budget.UpdateMonthlyBudget)
	api.DELETE("/monthly-budgets/:id", budget.DeleteMonthlyBudget)

	// YearlyBudget endpoints
	api.POST("/yearly-budgets", budget.CreateYearlyBudget)
	api.GET("/yearly-budgets/:id", budget.GetYearlyBudgetByID)
	api.GET("/yearly-budgets", budget.ListYearlyBudgets)
	api.PUT("/yearly-budgets/:id", budget.UpdateYearlyBudget)
	api.DELETE("/yearly-budgets/:id", budget.DeleteYearlyBudget)

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
