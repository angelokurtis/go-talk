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
		return errors.Errorf("failed to create manager: %w", err)
	}

	defer cleanup()

	slog.InfoContext(ctx, "Manager created")

	speech, err := mgr.SpeechGenerator.GenerateSpeech(ctx, "Defesa que ninguém passa\nLinha atacante de raça\nTorcida que canta e vibra\n\nPor nosso Alviverde inteiro\nQue sabe ser brasileiro\nOstentando a sua fibra\n\n")
	if err != nil {
		return err
	}

	return mgr.MP3Writer.Write(ctx, speech)
}
