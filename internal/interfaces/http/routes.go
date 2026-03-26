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
}
