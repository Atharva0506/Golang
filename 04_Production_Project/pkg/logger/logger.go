package logger

import (
	"log/slog"
	"os"

	"github.com/Atharva0506/trading_bot/internal/config"
)

func NewLogger(cfg *config.LoggerConfig) *slog.Logger {
	var slogLevel slog.Level
	switch cfg.Level {
	case "debug":
		slogLevel = slog.LevelDebug
	case "info":
		slogLevel = slog.LevelInfo
	case "warn":
		slogLevel = slog.LevelWarn
	case "error":
		slogLevel = slog.LevelError
	default:
		slogLevel = slog.LevelInfo
	}
	opts := &slog.HandlerOptions{
		Level: slogLevel,
	}
	var handler slog.Handler

	if cfg.Format == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}
