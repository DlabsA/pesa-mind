package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGamificationEndpoints(t *testing.T) {
	baseURL := "http://localhost:8080/api/v1"
	jwt := os.Getenv("TEST_JWT")

	t.Run("List Badges", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/gamification/badges", nil)
		req.Header.Set("Authorization", "Bearer "+jwt)
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var badges []map[string]interface{}
		_ = json.NewDecoder(resp.Body).Decode(&badges)
		_ = resp.Body.Close()
	})

	t.Run("List Streaks", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/gamification/streaks", nil)
		req.Header.Set("Authorization", "Bearer "+jwt)
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		_ = resp.Body.Close()
	})

	t.Run("List Achievements", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/gamification/achievements", nil)
		req.Header.Set("Authorization", "Bearer "+jwt)
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		_ = resp.Body.Close()
	})

	t.Run("Get Leaderboard", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/gamification/leaderboard", nil)
		req.Header.Set("Authorization", "Bearer "+jwt)
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		_ = resp.Body.Close()
	})

	t.Run("List Rewards", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/gamification/rewards", nil)
		req.Header.Set("Authorization", "Bearer "+jwt)
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		_ = resp.Body.Close()
	})

	t.Run("Award First Transaction Badge", func(t *testing.T) {
		// Create a transaction
		body := `{"account_id":"some-account-uuid","category_id":"some-category-uuid","amount":10.0,"tx_type":"expense","note":"test","date":` + "`date`" + `}`
		req, _ := http.NewRequest("POST", baseURL+"/transactions", bytes.NewBufferString(body))
		req.Header.Set("Authorization", "Bearer "+jwt)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		_ = resp.Body.Close()
		// List badges
		req2, _ := http.NewRequest("GET", baseURL+"/gamification/badges", nil)
		req2.Header.Set("Authorization", "Bearer "+jwt)
		resp2, err := http.DefaultClient.Do(req2)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp2.StatusCode)
		var badges []map[string]interface{}
		_ = json.NewDecoder(resp2.Body).Decode(&badges)
		_ = resp2.Body.Close()
		found := false
		for _, b := range badges {
			if b["id"] == "11111111-1111-1111-1111-111111111111" {
				found = true
			}
		}
		assert.True(t, found, "First Transaction badge should be awarded")
	})

	t.Run("Award Goal Achiever Achievement", func(t *testing.T) {
		// Simulate goal completion (assume endpoint exists)
		goalID := "some-goal-uuid"
		req, _ := http.NewRequest("POST", baseURL+"/savingsgoals/"+goalID+"/complete", nil)
		req.Header.Set("Authorization", "Bearer "+jwt)
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		_ = resp.Body.Close()
		// List achievements
		req2, _ := http.NewRequest("GET", baseURL+"/gamification/achievements", nil)
		req2.Header.Set("Authorization", "Bearer "+jwt)
		resp2, err := http.DefaultClient.Do(req2)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp2.StatusCode)
		var achievements []map[string]interface{}
		_ = json.NewDecoder(resp2.Body).Decode(&achievements)
		_ = resp2.Body.Close()
		found := false
		for _, a := range achievements {
			if a["id"] == "22222222-2222-2222-2222-222222222222" {
				found = true
			}
		}
		assert.True(t, found, "Goal Achiever achievement should be awarded")
	})
}
