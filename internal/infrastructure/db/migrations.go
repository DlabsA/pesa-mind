package db

import (
	"log"

	"pesa-mind/internal/domain/account"
	"pesa-mind/internal/domain/analytics"
	"pesa-mind/internal/domain/automation"
	"pesa-mind/internal/domain/budget"
	"pesa-mind/internal/domain/category"
	"pesa-mind/internal/domain/gamification"
	"pesa-mind/internal/domain/notification"
	"pesa-mind/internal/domain/savingsgoal"
	"pesa-mind/internal/domain/transaction"
	"pesa-mind/internal/domain/user"

	"gorm.io/gorm"
)

// AutoMigrateAll runs all domain migrations
func AutoMigrateAll(database *gorm.DB) error {
	// Core domain tables
	if err := database.AutoMigrate(
		&user.User{},
		&user.Profile{},
		&account.Account{},
		&category.Category{},
		&transaction.Transaction{},
	); err != nil {
		log.Fatalf("failed to migrate core tables: %v", err)
		return err
	}

	// Budget tables
	if err := database.AutoMigrate(&budget.Budget{}); err != nil {
		log.Fatalf("failed to migrate budget tables: %v", err)
		return err
	}

	// Savings Goal tables
	if err := database.AutoMigrate(&savingsgoal.SavingsGoal{}); err != nil {
		log.Fatalf("failed to migrate savings goal tables: %v", err)
		return err
	}

	// Analytics tables
	if err := database.AutoMigrate(&analytics.AnalyticsSnapshot{}); err != nil {
		log.Fatalf("failed to migrate analytics tables: %v", err)
		return err
	}

	// Automation tables
	if err := database.AutoMigrate(&automation.SMSAutomation{}); err != nil {
		log.Fatalf("failed to migrate automation tables: %v", err)
		return err
	}

	// Notification tables
	if err := database.AutoMigrate(
		&notification.Notification{},
		&notification.Preference{},
	); err != nil {
		log.Fatalf("failed to migrate notification tables: %v", err)
		return err
	}

	// Gamification tables
	if err := database.AutoMigrate(
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
		return err
	}

	return nil
}
