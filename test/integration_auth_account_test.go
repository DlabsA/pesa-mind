//go:build integration

package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type registerReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type loginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type accountReq struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Currency string `json:"currency"`
}
type accountResp struct {
	ID string `json:"id"`
}

func TestRegisterLoginCreateAccount(t *testing.T) {
	baseURL := os.Getenv("INTEGRATION_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080/api/v1"
	}
	// Register
	reg := registerReq{Email: "testuser@example.com", Password: "testpass123"}
	regBody, _ := json.Marshal(reg)
	resp, err := http.Post(baseURL+"/users/register", "application/json", bytes.NewReader(regBody))
	assert.NoError(t, err)
	if err != nil {
		t.Fatalf("Failed to POST /users/register: %v", err)
	}
	assert.NotNil(t, resp)
	if resp == nil {
		t.Fatalf("No response from /users/register")
	}
	assert.Equal(t, 201, resp.StatusCode)
	// Login
	login := loginReq{Email: reg.Email, Password: reg.Password}
	loginBody, _ := json.Marshal(login)
	resp, err = http.Post(baseURL+"/auth/login", "application/json", bytes.NewReader(loginBody))
	assert.NoError(t, err)
	if err != nil {
		t.Fatalf("Failed to POST /auth/login: %v", err)
	}
	assert.NotNil(t, resp)
	if resp == nil {
		t.Fatalf("No response from /auth/login")
	}
	assert.Equal(t, 200, resp.StatusCode)
	var lresp loginResp
	json.NewDecoder(resp.Body).Decode(&lresp)
	assert.NotEmpty(t, lresp.AccessToken)
	// Create Account
	acc := accountReq{Name: "Wallet", Type: "cash", Currency: "USD"}
	accBody, _ := json.Marshal(acc)
	req, _ := http.NewRequest("POST", baseURL+"/accounts", bytes.NewReader(accBody))
	req.Header.Set("Authorization", "Bearer "+lresp.AccessToken)
	client := &http.Client{}
	resp3, err := client.Do(req)
	assert.NoError(t, err)
	if err != nil {
		t.Fatalf("Failed to POST /accounts: %v", err)
	}
	assert.NotNil(t, resp3)
	if resp3 == nil {
		t.Fatalf("No response from /accounts")
	}
	assert.Equal(t, 201, resp3.StatusCode)
	var aresp accountResp
	json.NewDecoder(resp3.Body).Decode(&aresp)
	assert.NotEmpty(t, aresp.ID)
}
