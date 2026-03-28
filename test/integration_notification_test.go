package test

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNotificationEndpoints(t *testing.T) {
	// Start the server (assume main.go runs on :8080)
	baseURL := "http://localhost:8080/api/v1"
	jwt := os.Getenv("TEST_JWT") // Set a valid JWT for testing

	t.Run("List Notifications", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/notifications", nil)
		req.Header.Set("Authorization", "Bearer "+jwt)
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var notifs []map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&notifs)
		assert.NoError(t, err)
		err = resp.Body.Close()
		assert.NoError(t, err)
	})

	t.Run("Mark Notification As Read", func(t *testing.T) {
		// First, list notifications to get an ID
		req, _ := http.NewRequest("GET", baseURL+"/notifications", nil)
		req.Header.Set("Authorization", "Bearer "+jwt)
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		var notifs []map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&notifs)
		assert.NoError(t, err)
		err = resp.Body.Close()
		assert.NoError(t, err)
		if len(notifs) == 0 {
			t.Skip("No notifications to mark as read")
		}
		id := notifs[0]["id"].(string)
		// Mark as read
		req2, _ := http.NewRequest("POST", baseURL+"/notifications/"+id+"/read", nil)
		req2.Header.Set("Authorization", "Bearer "+jwt)
		resp2, err := http.DefaultClient.Do(req2)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp2.StatusCode)
		err = resp2.Body.Close()
		assert.NoError(t, err)
	})

	t.Run("Get Notification Preferences", func(t *testing.T) {
		req, _ := http.NewRequest("GET", baseURL+"/notifications/preferences", nil)
		req.Header.Set("Authorization", "Bearer "+jwt)
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		err = resp.Body.Close()
		assert.NoError(t, err)
	})

	t.Run("Set Notification Preferences", func(t *testing.T) {
		body := `{"in_app":true,"push":true,"email":false}`
		req, _ := http.NewRequest("POST", baseURL+"/notifications/preferences",
			strings.NewReader(body))
		req.Header.Set("Authorization", "Bearer "+jwt)
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		err = resp.Body.Close()
		assert.NoError(t, err)
	})

	// Optionally: test push/email stub endpoints if exposed

	time.Sleep(100 * time.Millisecond) // Allow async DB writes to settle
}
