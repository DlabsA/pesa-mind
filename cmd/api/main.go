package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"pesa-mind/internal/config"
	"pesa-mind/internal/infrastructure/db"
	"pesa-mind/internal/infrastructure/logger"

	// Domain modules
	"pesa-mind/internal/domain/account"
	"pesa-mind/internal/domain/analytics"
	"pesa-mind/internal/domain/budget"
	"pesa-mind/internal/domain/category"
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

	// GORM auto-migrate
	db.DB.AutoMigrate(&user.User{}, &account.Account{}, &category.Category{}, &transaction.Transaction{})

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

	// Instantiate services
	userService := user.NewService(userRepo)
	accountService := account.NewService(accountRepo)
	categoryService := category.NewService(categoryRepo)
	transactionService := transaction.NewService(transactionRepo)
	budgetService := budget.NewService(budgetRepo)
	savingsGoalService := savingsgoal.NewService(savingsGoalRepo)
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
		}
	}

	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	if err := engine.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
