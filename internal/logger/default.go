package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

// SetUp initializes the logger with predefined settings and returns the logger instance.
func SetUp() *slog.Logger {
	// Create a new handler with tint colorized output
	handler := tint.NewHandler(os.Stderr, &tint.Options{
		AddSource:  true,
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
	})

	// Create a new slog logger instance with the formatted handler
	l := slog.New(handler)

	// Set the newly created logger as the default logger for the application
	slog.SetDefault(l)

	// Return the logger instance for potential further configuration
	return l
}
