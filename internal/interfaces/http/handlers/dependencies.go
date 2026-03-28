package handlers

import (
	"pesa-mind/internal/domain/analytics"
	"pesa-mind/internal/domain/budget"
	"pesa-mind/internal/domain/gamification"
	// ...existing imports...
)

type Dependencies struct {
	BudgetService *budget.Service
	// ...existing fields...
	AnalyticsServiceCRUD *analytics.Service
	AnalyticsService     *analytics.AnalyticsService
	GamificationService  *gamification.Service
}
