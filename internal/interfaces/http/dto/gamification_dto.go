package dto

type BadgeResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	EarnedAt    int64  `json:"earned_at,omitempty"`
}

type StreakResponse struct {
	Type   string `json:"type"`
	Count  int    `json:"count"`
	Active bool   `json:"active"`
}

type AchievementResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	EarnedAt    int64  `json:"earned_at,omitempty"`
}

type LeaderboardEntryResponse struct {
	UserID string `json:"user_id"`
	Score  int    `json:"score"`
	Rank   int    `json:"rank"`
}

type RewardResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Points      int    `json:"points"`
	ClaimedAt   int64  `json:"claimed_at,omitempty"`
}
