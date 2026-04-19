package setup

import (
	"pesa-mind/internal/config"
	"pesa-mind/internal/domain/analytics"
	"pesa-mind/internal/domain/automation"
	"pesa-mind/internal/domain/budget"
	"pesa-mind/internal/domain/category"
	"pesa-mind/internal/domain/gamification"
	"pesa-mind/internal/domain/notification"
	"pesa-mind/internal/domain/savingsgoal"
	"pesa-mind/internal/domain/transaction"
	"pesa-mind/internal/domain/user"
	"pesa-mind/internal/interfaces/http/handlers"

	"gorm.io/gorm"
)

// AppDependencies holds all domain handlers organized by module
type AppDependencies struct {
	UserHandler          *handlers.UserHandler
	AuthHandler          *handlers.AuthHandler
	ChannelDetailHandler *handlers.CategoryHandler
	TransactionHandler   *handlers.TransactionHandler
	BudgetHandler        *handlers.BudgetHandler
	SavingsGoalHandler   *handlers.SavingsGoalHandler
	AnalyticsHandler     *handlers.AnalyticsHandler
	AutomationHandler    *handlers.AutomationHandler
	AutomationSMSHandler *handlers.AutomationSMSHandler
	NotificationHandler  *handlers.NotificationHandler
	PreferenceHandler    *handlers.NotificationPreferenceHandler
	GamificationHandler  *handlers.GamificationHandler
}

// Initialize sets up all domain modules in dependency order
func Initialize(cfg *config.Config, database *gorm.DB) *AppDependencies {
	// Initialize independent domains first (no dependencies)
	userRepo := user.NewGormUserRepository(database)
	userService := user.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	channelDetailRepo := category.NewGormChannelDetailsRepository(database)
	channelDetailService := category.NewService(channelDetailRepo)
	channelDetailHandler := handlers.NewCategoryHandler(channelDetailService)

	budgetRepo := budget.NewGormRepository(database)
	budgetService := budget.NewService(budgetRepo)
	budgetHandler := handlers.NewBudgetHandler(budgetService)

	// Initialize gamification (no business logic dependencies)
	gamificationRepo := gamification.NewGormRepository(database)
	gamificationService := gamification.NewService(gamificationRepo)
	gamificationHandler := handlers.NewGamificationHandler(gamificationService)

	// Initialize domains that depend on gamification
	transactionRepo := transaction.NewGormTransactionRepository(database)
	transactionService := transaction.NewService(transactionRepo, gamificationService, userService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	savingsGoalRepo := savingsgoal.NewGormRepository(database)
	savingsGoalService := savingsgoal.NewService(savingsGoalRepo, gamificationService)
	savingsGoalHandler := handlers.NewSavingsGoalHandler(savingsGoalService)

	// Initialize auth (depends on user)
	authHandler := handlers.NewAuthHandler(userService, cfg.JWTSecret)

	// Initialize analytics (depends on repositories)
	analyticsRepo := analytics.NewGormRepository(database)
	analyticsService := analytics.NewService(analyticsRepo)
	analyticsServiceRT := analytics.NewAnalyticsService(database, transactionRepo, budgetRepo, savingsGoalRepo)
	analyticsHandler := handlers.NewAnalyticsHandler(analyticsService, analyticsServiceRT)

	// Initialize automation (depends on transaction service)
	automationRepo := automation.NewGormRepository(database)
	automationService := automation.NewService(automationRepo)
	automationHandler := handlers.NewAutomationHandler(automationService)
	automationSMSHandler := handlers.NewAutomationSMSHandler(transactionService)

	// Initialize notification
	notificationRepo := notification.NewGormRepository(database)
	notificationService := notification.NewService(notificationRepo)
	notificationHandler := handlers.NewNotificationHandler(notificationService)

	preferenceRepo := notification.NewGormPreferenceRepository(database)
	preferenceService := notification.NewPreferenceService(preferenceRepo)
	preferenceHandler := handlers.NewNotificationPreferenceHandler(preferenceService)

	return &AppDependencies{
		UserHandler:          userHandler,
		AuthHandler:          authHandler,
		ChannelDetailHandler: channelDetailHandler,
		TransactionHandler:   transactionHandler,
		BudgetHandler:        budgetHandler,
		SavingsGoalHandler:   savingsGoalHandler,
		AnalyticsHandler:     analyticsHandler,
		AutomationHandler:    automationHandler,
		AutomationSMSHandler: automationSMSHandler,
		NotificationHandler:  notificationHandler,
		PreferenceHandler:    preferenceHandler,
		GamificationHandler:  gamificationHandler,
	}
}
