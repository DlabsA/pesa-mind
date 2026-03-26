package dto

import (
	"encoding/json"
	"testing"
)

func TestAccountResponseJSON(t *testing.T) {
	resp := AccountResponse{ID: "id", UserID: "uid", Name: "Wallet", Type: "cash", Currency: "USD", Balance: 100}
	_, err := json.Marshal(resp)
	if err != nil {
		t.Errorf("Failed to marshal AccountResponse: %v", err)
	}
}

func TestCreateAccountRequestJSON(t *testing.T) {
	jsonStr := `{"name":"Wallet","type":"cash","currency":"USD"}`
	var req CreateAccountRequest
	err := json.Unmarshal([]byte(jsonStr), &req)
	if err != nil {
		t.Errorf("Failed to unmarshal CreateAccountRequest: %v", err)
	}
	if req.Name != "Wallet" || req.Type != "cash" || req.Currency != "USD" {
		t.Errorf("Unexpected values: %+v", req)
	}
}

func TestUserResponseJSON(t *testing.T) {
	resp := UserResponse{ID: "id", Email: "user@example.com"}
	_, err := json.Marshal(resp)
	if err != nil {
		t.Errorf("Failed to marshal UserResponse: %v", err)
	}
}

func TestRegisterRequestJSON(t *testing.T) {
	jsonStr := `{"email":"user@example.com","password":"password123"}`
	var req RegisterRequest
	err := json.Unmarshal([]byte(jsonStr), &req)
	if err != nil {
		t.Errorf("Failed to unmarshal RegisterRequest: %v", err)
	}
	if req.Email != "user@example.com" || req.Password != "password123" {
		t.Errorf("Unexpected values: %+v", req)
	}
}

func TestCategoryResponseJSON(t *testing.T) {
	resp := CategoryResponse{ID: "id", UserID: "uid", Name: "Food", Type: "expense", ParentID: nil}
	_, err := json.Marshal(resp)
	if err != nil {
		t.Errorf("Failed to marshal CategoryResponse: %v", err)
	}
}

func TestCreateCategoryRequestJSON(t *testing.T) {
	jsonStr := `{"name":"Food","type":"expense"}`
	var req CreateCategoryRequest
	err := json.Unmarshal([]byte(jsonStr), &req)
	if err != nil {
		t.Errorf("Failed to unmarshal CreateCategoryRequest: %v", err)
	}
	if req.Name != "Food" || req.Type != "expense" {
		t.Errorf("Unexpected values: %+v", req)
	}
}

func TestTransactionResponseJSON(t *testing.T) {
	resp := TransactionResponse{ID: "id", UserID: "uid", AccountID: "aid", CategoryID: "cid", Amount: 10.5, Type: "expense", Note: "Lunch", Date: 1234567890}
	_, err := json.Marshal(resp)
	if err != nil {
		t.Errorf("Failed to marshal TransactionResponse: %v", err)
	}
}

func TestCreateTransactionRequestJSON(t *testing.T) {
	jsonStr := `{"account_id":"aid","category_id":"cid","amount":10.5,"type":"expense","note":"Lunch","date":1234567890}`
	var req CreateTransactionRequest
	err := json.Unmarshal([]byte(jsonStr), &req)
	if err != nil {
		t.Errorf("Failed to unmarshal CreateTransactionRequest: %v", err)
	}
	if req.AccountID != "aid" || req.CategoryID != "cid" || req.Amount != 10.5 || req.Type != "expense" || req.Note != "Lunch" || req.Date != 1234567890 {
		t.Errorf("Unexpected values: %+v", req)
	}
}

func TestLoginResponseJSON(t *testing.T) {
	resp := LoginResponse{AccessToken: "token", RefreshToken: "refresh"}
	_, err := json.Marshal(resp)
	if err != nil {
		t.Errorf("Failed to marshal LoginResponse: %v", err)
	}
}

func TestLoginRequestJSON(t *testing.T) {
	jsonStr := `{"email":"user@example.com","password":"password123"}`
	var req LoginRequest
	err := json.Unmarshal([]byte(jsonStr), &req)
	if err != nil {
		t.Errorf("Failed to unmarshal LoginRequest: %v", err)
	}
	if req.Email != "user@example.com" || req.Password != "password123" {
		t.Errorf("Unexpected values: %+v", req)
	}
}

func TestRefreshRequestJSON(t *testing.T) {
	jsonStr := `{"refresh_token":"refresh"}`
	var req RefreshRequest
	err := json.Unmarshal([]byte(jsonStr), &req)
	if err != nil {
		t.Errorf("Failed to unmarshal RefreshRequest: %v", err)
	}
	if req.RefreshToken != "refresh" {
		t.Errorf("Unexpected values: %+v", req)
	}
}

func TestRefreshResponseJSON(t *testing.T) {
	resp := RefreshResponse{AccessToken: "token", RefreshToken: "refresh"}
	_, err := json.Marshal(resp)
	if err != nil {
		t.Errorf("Failed to marshal RefreshResponse: %v", err)
	}
}
