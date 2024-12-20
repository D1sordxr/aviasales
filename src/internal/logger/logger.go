package logger

import (
	"github.com/D1sordxr/aviasales/src/internal/config/config"
	"github.com/D1sordxr/aviasales/src/internal/logger/handlers/designed"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Logger struct {
	*slog.Logger
}

func NewLogger(config *config.Config) *Logger {
	var logger *slog.Logger
	var handler slog.Handler

	switch config.AppConfig.Mode {
	case envLocal:
		logger = designed.NewPrettySlog()
		return &Logger{logger}
	case envDev:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case envProd:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	}

	logger = slog.New(handler)

	return &Logger{logger}
}
