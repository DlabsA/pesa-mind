package model

// Re-export all domain models so they can be referenced as model.Transaction, model.User, etc.

import (
	"pesa-mind/internal/domain/analytics"
	"pesa-mind/internal/domain/automation"
	"pesa-mind/internal/domain/budget"
	"pesa-mind/internal/domain/category"
	"pesa-mind/internal/domain/gamification"
	"pesa-mind/internal/domain/notification"
	"pesa-mind/internal/domain/savingsgoal"
	"pesa-mind/internal/domain/transaction"
	"pesa-mind/internal/domain/user"
)

// Transaction models
type Transaction = transaction.Transaction

// User models
type User = user.User
type Profile = user.Profile

// Category models
type Category = category.Category

// Budget models
type Budget = budget.Budget

// SavingsGoal models
type SavingsGoal = savingsgoal.SavingsGoal

// Gamification models
type Badge = gamification.Badge
type Achievement = gamification.Achievement
type Streak = gamification.Streak

// Automation models
type SMSAutomation = automation.SMSAutomation

// Analytics models
type AnalyticsSnapshot = analytics.AnalyticsSnapshot

// Notification models
type Notification = notification.Notification
type Preference = notification.Preference
