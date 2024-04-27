package main

import (
	"context"
	"log/slog"

	"github.com/lmittmann/tint"

	"github.com/angelokurtis/go-talk/internal/errors"
	"github.com/angelokurtis/go-talk/internal/logger"
)

func main() {
	// Creating a context for the main function
	ctx := context.Background()

	// Set up logging
	l := logger.SetUp()

	// Executes the main logic of the program
	err := run(ctx)
	if err == nil {
		return // Exiting the main function if there was no error
	}

	// Handling errors
	switch t := err.(type) {
	case errors.Traceable:
		l.ErrorContext(ctx, t.String(), tint.Err(err))
	default:
		l.ErrorContext(ctx, "Failed", tint.Err(err))
	}
}

func run(ctx context.Context) error {
	mgr, cleanup, err := NewManager()
	if err != nil {
		return err
	}

	defer cleanup()

	msg, err := mgr.Example.Show()
	if err != nil {
		return err
	}

	slog.InfoContext(ctx, msg)

	return nil
}
