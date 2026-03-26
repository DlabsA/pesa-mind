package db

import (
	"os"
	"testing"
)

func TestInitDBConnection(t *testing.T) {
	if os.Getenv("DB_HOST") == "" {
		t.Skip("DB_HOST not set; skipping DB connection test")
	}
	err := Init()
	if err != nil {
		t.Errorf("Init() failed: %v", err)
	}
	if DB == nil {
		t.Error("DB global variable is nil after Init()")
	}
}
