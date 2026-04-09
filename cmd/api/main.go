//go:generate autodoc-gen -router=gin

package main

import (
	"log"
	"pesa-mind/internal/config"
	"pesa-mind/internal/infrastructure/db"
	"pesa-mind/internal/infrastructure/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// Domain modules
	"pesa-mind/internal/domain/account"
	"pesa-mind/internal/domain/analytics"
	"pesa-mind/internal/domain/budget"
	"pesa-mind/internal/domain/category"
	"pesa-mind/internal/domain/gamification"
	"pesa-mind/internal/domain/savingsgoal"
	"pesa-mind/internal/domain/transaction"
	"pesa-mind/internal/domain/user"

	// Handlers
	handlers "pesa-mind/internal/interfaces/http/handlers"
	"pesa-mind/internal/interfaces/http/middleware"

	// Automation
	"pesa-mind/internal/domain/automation"

	// Notifications
	"pesa-mind/internal/domain/notification"
)

func main() {
	_ = godotenv.Load()
	cfg := config.LoadConfig()
	logger.Init(cfg.LogLevel)
	if err := db.Init(); err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	// GORM auto-migrate - ALL domain models
	// Core domain tables
	if err := db.DB.AutoMigrate(
		&user.User{},
		&user.Profile{},
		&account.Account{},
		&category.Category{},
		&transaction.Transaction{},
	); err != nil {
		log.Fatalf("failed to migrate core tables: %v", err)
	}

	// Budget tables
	if err := db.DB.AutoMigrate(&budget.Budget{}); err != nil {
		log.Fatalf("failed to migrate budget tables: %v", err)
	}

	// Savings Goal tables
	if err := db.DB.AutoMigrate(&savingsgoal.SavingsGoal{}); err != nil {
		log.Fatalf("failed to migrate savings goal tables: %v", err)
	}

	// Analytics tables
	if err := db.DB.AutoMigrate(&analytics.AnalyticsSnapshot{}); err != nil {
		log.Fatalf("failed to migrate analytics tables: %v", err)
	}

	// Automation tables
	if err := db.DB.AutoMigrate(&automation.SMSAutomation{}); err != nil {
		log.Fatalf("failed to migrate automation tables: %v", err)
	}

	// Notification tables
	if err := db.DB.AutoMigrate(
		&notification.Notification{},
		&notification.Preference{},
	); err != nil {
		log.Fatalf("failed to migrate notification tables: %v", err)
	}

	// Gamification tables
	if err := db.DB.AutoMigrate(
		&gamification.Badge{},
		&gamification.UserBadge{},
		&gamification.Streak{},
		&gamification.Achievement{},
		&gamification.UserAchievement{},
		&gamification.LeaderboardEntry{},
		&gamification.Reward{},
		&gamification.UserReward{},
	); err != nil {
		log.Fatalf("failed to migrate gamification tables: %v", err)
	}

	// Instantiate repositories
	userRepo := user.NewGormUserRepository(db.DB)
	accountRepo := account.NewGormAccountRepository(db.DB)
	categoryRepo := category.NewGormCategoryRepository(db.DB)
	transactionRepo := transaction.NewGormTransactionRepository(db.DB)
	budgetRepo := budget.NewGormRepository(db.DB)
	savingsGoalRepo := savingsgoal.NewGormRepository(db.DB)
	analyticsRepo := analytics.NewGormRepository(db.DB)

	// Automation module
	automationRepo := automation.NewGormRepository(db.DB)

	// Notification module
	notificationRepo := notification.NewGormRepository(db.DB)

	// Gamification module
	gamificationRepo := gamification.NewGormRepository(db.DB)
	gamificationService := gamification.NewService(gamificationRepo)

	// Instantiate services
	userService := user.NewService(userRepo)
	accountService := account.NewService(accountRepo)
	categoryService := category.NewService(categoryRepo)
	transactionService := transaction.NewService(transactionRepo, gamificationService)
	budgetService := budget.NewService(budgetRepo)
	savingsGoalService := savingsgoal.NewService(savingsGoalRepo, gamificationService)
	analyticsService := analytics.NewService(analyticsRepo)
	analyticsServiceRealtime := analytics.NewAnalyticsService(db.DB, transactionRepo, budgetRepo, savingsGoalRepo)
	automationService := automation.NewService(automationRepo)
	notificationService := notification.NewService(notificationRepo)

	// Notification preferences
	preferenceRepo := notification.NewGormPreferenceRepository(db.DB)
	preferenceService := notification.NewPreferenceService(preferenceRepo)

	// Instantiate handlers
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(userService, cfg.JWTSecret)
	accountHandler := handlers.NewAccountHandler(accountService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	budgetHandler := handlers.NewBudgetHandler(budgetService)
	savingsGoalHandler := handlers.NewSavingsGoalHandler(savingsGoalService)
	analyticsHandler := handlers.NewAnalyticsHandler(analyticsService, analyticsServiceRealtime)
	automationHandler := handlers.NewAutomationHandler(automationService)
	automationSMSHandler := handlers.NewAutomationSMSHandler(transactionService)
	notificationHandler := handlers.NewNotificationHandler(notificationService)
	preferenceHandler := handlers.NewNotificationPreferenceHandler(preferenceService)
	gamificationHandler := handlers.NewGamificationHandler(gamificationService)

	engine := gin.Default()

	// Register routes
	api := engine.Group("/api/v1")
	{
		api.POST("/users/register", userHandler.Register)
		api.POST("/auth/login", authHandler.Login)
		api.POST("/auth/refresh", authHandler.Refresh)

		// Protected routes
		auth := api.Group("")
		auth.Use(middleware.JWTAuthMiddleware())
		{
			auth.POST("/accounts", accountHandler.Create)
			auth.GET("/accounts", accountHandler.List)
			auth.POST("/categories", categoryHandler.Create)
			auth.GET("/categories", categoryHandler.List)
			auth.POST("/transactions", transactionHandler.Create)
			auth.GET("/transactions", transactionHandler.List)
			auth.POST("/budgets", budgetHandler.Create)
			auth.GET("/budgets", budgetHandler.List)
			auth.POST("/savingsgoals", savingsGoalHandler.Create)
			auth.GET("/savingsgoals", savingsGoalHandler.List)
			auth.POST("/analytics", analyticsHandler.Create)
			auth.GET("/analytics", analyticsHandler.List)
			auth.GET("/analytics/income", analyticsHandler.TotalIncome)
			auth.GET("/analytics/expenses", analyticsHandler.TotalExpenses)
			auth.GET("/analytics/budget-utilization", analyticsHandler.BudgetUtilization)
			auth.GET("/analytics/savings-progress", analyticsHandler.SavingsProgress)
			// Automation endpoints
			auth.POST("/automation/sms", automationHandler.CreateRule)
			auth.GET("/automation/sms", automationHandler.ListRules)
			// Automation SMS transaction endpoint
			auth.POST("/automation/sms/transaction", automationSMSHandler.CreateTransactionFromSMS)
			// Notification endpoints
			auth.GET("/notifications", notificationHandler.List)
			auth.POST("/notifications/:id/read", notificationHandler.MarkAsRead)
			// Notification preference endpoints
			auth.GET("/notifications/preferences", preferenceHandler.Get)
			auth.POST("/notifications/preferences", preferenceHandler.Set)

			// Gamification endpoints
			auth.GET("/gamification/badges", gamificationHandler.ListBadges)
			auth.GET("/gamification/streaks", gamificationHandler.ListStreaks)
			auth.GET("/gamification/achievements", gamificationHandler.ListAchievements)
			auth.GET("/gamification/leaderboard", gamificationHandler.GetLeaderboard)
			auth.GET("/gamification/rewards", gamificationHandler.ListRewards)
			auth.POST("/gamification/rewards/:reward_id/claim", gamificationHandler.ClaimReward)
		}
	}

	// Add a health check endpoint
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	if err := engine.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
