package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // minimum level logged
	})
	return slog.New(handler)
}
