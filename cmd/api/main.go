//go:generate autodoc-gen -router=gin

package main

import (
	"log"
	"pesa-mind/internal/config"
	"pesa-mind/internal/infrastructure/db"
	"pesa-mind/internal/infrastructure/logger"
	"pesa-mind/internal/infrastructure/setup"
	"pesa-mind/internal/interfaces/http/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// AppDependencies is a type alias for cleaner code
type AppDependencies = setup.AppDependencies

func main() {
	// Load environment
	_ = godotenv.Load()
	cfg := config.LoadConfig()
	logger.Init(cfg.LogLevel)

	// Initialize database
	if err := db.Init(); err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	// Run migrations
	if err := db.AutoMigrateAll(db.DB); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	// Initialize all domain dependencies
	deps := setup.Initialize(cfg, db.DB)

	// Create Gin engine and register routes
	engine := gin.Default()
	registerRoutes(engine, deps)

	// Health check endpoint
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Start server
	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	if err := engine.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// registerRoutes registers all domain routes
func registerRoutes(engine *gin.Engine, deps *AppDependencies) {
	api := engine.Group("/api/v1")
	{
		// Public routes
		api.POST("/users/register", deps.UserHandler.Register)
		api.POST("/auth/login", deps.AuthHandler.Login)
		api.POST("/auth/refresh", deps.AuthHandler.Refresh)

		// Protected user routes
		userAuth := api.Group("/users/me")
		userAuth.Use(middleware.JWTAuthMiddleware())
		{
			userAuth.GET("", deps.UserHandler.Get)
			userAuth.PATCH("", deps.UserHandler.Update)
			userAuth.POST("/change_password", deps.UserHandler.ChangePassword)
		}

		// Protected routes (require authentication)
		auth := api.Group("")
		auth.Use(middleware.JWTAuthMiddleware())
		{

			// Category routes
			auth.POST("/categories", deps.ChannelDetailHandler.Create)
			auth.GET("/categories", deps.ChannelDetailHandler.List)
			auth.GET("/categories/channel-type", deps.ChannelDetailHandler.GetByChannelType)
			auth.GET("/categories/status", deps.ChannelDetailHandler.GetByStatus)
			auth.PATCH("/categories/:id", deps.ChannelDetailHandler.Update)
			auth.DELETE("/categories/:id", deps.ChannelDetailHandler.Delete)

			// Transaction routes
			auth.POST("/transactions", deps.TransactionHandler.Create)
			auth.GET("/transactions", deps.TransactionHandler.List)

			// Budget routes
			auth.POST("/budgets", deps.BudgetHandler.Create)
			auth.GET("/budgets", deps.BudgetHandler.List)

			// Savings Goal routes
			auth.POST("/savingsgoals", deps.SavingsGoalHandler.Create)
			auth.GET("/savingsgoals", deps.SavingsGoalHandler.List)

			// Analytics routes
			auth.POST("/analytics", deps.AnalyticsHandler.Create)
			auth.GET("/analytics", deps.AnalyticsHandler.List)
			auth.GET("/analytics/income", deps.AnalyticsHandler.TotalIncome)
			auth.GET("/analytics/expenses", deps.AnalyticsHandler.TotalExpenses)
			auth.GET("/analytics/budget-utilization", deps.AnalyticsHandler.BudgetUtilization)
			auth.GET("/analytics/savings-progress", deps.AnalyticsHandler.SavingsProgress)

			// Automation routes
			auth.POST("/automation/sms", deps.AutomationHandler.CreateRule)
			auth.GET("/automation/sms", deps.AutomationHandler.ListRules)
			auth.POST("/automation/sms/transaction", deps.AutomationSMSHandler.CreateTransactionFromSMS)

			// Notification routes
			auth.GET("/notifications", deps.NotificationHandler.List)
			auth.POST("/notifications/:id/read", deps.NotificationHandler.MarkAsRead)

			// Notification preference routes
			auth.GET("/notifications/preferences", deps.PreferenceHandler.Get)
			auth.POST("/notifications/preferences", deps.PreferenceHandler.Set)

			// Gamification routes
			auth.GET("/gamification/badges", deps.GamificationHandler.ListBadges)
			auth.GET("/gamification/streaks", deps.GamificationHandler.ListStreaks)
			auth.GET("/gamification/achievements", deps.GamificationHandler.ListAchievements)
			auth.GET("/gamification/leaderboard", deps.GamificationHandler.GetLeaderboard)
			auth.GET("/gamification/rewards", deps.GamificationHandler.ListRewards)
			auth.POST("/gamification/rewards/:reward_id/claim", deps.GamificationHandler.ClaimReward)
		}
	}
}
