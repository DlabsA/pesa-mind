package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var Log zerolog.Logger

func Init(logLevel string) {
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		level = zerolog.InfoLevel
	}
	Log = zerolog.New(os.Stdout).Level(level).With().Timestamp().Logger()
}
