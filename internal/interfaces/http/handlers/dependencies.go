package handlers

import (
	"pesa-mind/internal/domain/analytics"
	"pesa-mind/internal/domain/budget"
	// ...existing imports...
)

type Dependencies struct {
	BudgetService *budget.Service
	// ...existing fields...
	AnalyticsServiceCRUD *analytics.Service
	AnalyticsService     *analytics.AnalyticsService
}
