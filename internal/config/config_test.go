package config

import (
	"os"
	"testing"
)

func TestLoadConfigDefaults(t *testing.T) {
	os.Clearenv()
	cfg := LoadConfig()
	if cfg.Port != "8080" && cfg.Port != "" {
		t.Errorf("Expected default port 8080 or empty, got %s", cfg.Port)
	}
	if cfg.DBHost == "" {
		t.Error("Expected DBHost to have a value")
	}
}
