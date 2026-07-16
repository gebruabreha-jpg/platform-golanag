// Package logger provides a structured, leveled logger built on the standard
// library's log/slog. It is intentionally generic so any service or CLI in
// this module (or a future one) can reuse it.
package logger

import (
	"log/slog"
	"os"
)

// New returns a structured JSON logger writing to stderr. JSON is preferred in
// containerized/production environments where logs are collected and parsed.
func New() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
}

// NewForEnv returns a logger appropriate for the given environment. "prod"
// (or empty) uses JSON; "dev"/"local" uses a human-readable text handler.
func NewForEnv(env string) *slog.Logger {
	if env == "dev" || env == "local" {
		return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	}
	return New()
}
