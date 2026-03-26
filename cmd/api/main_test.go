package main

import (
	"net/http"
	"testing"
)

func TestMainHealthCheck(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/api/v1/users/register")
	if err != nil {
		t.Skip("API server not running: ", err)
	}
	if resp.StatusCode != 405 && resp.StatusCode != 404 {
		t.Errorf("Expected 405 or 404 for GET /users/register, got %d", resp.StatusCode)
	}
}
