package logger

import (
	"testing"
)

func TestLoggerInit(t *testing.T) {
	Init("debug")
	// No panic or error expected
}
