package main

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/gotidy/ptr"
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

	body := azopenai.SpeechGenerationOptions{
		Input:          ptr.Of("Today is a wonderful day to build something people love!"),
		Voice:          ptr.Of(azopenai.SpeechVoiceAlloy),
		ResponseFormat: ptr.Of(azopenai.SpeechGenerationResponseFormatMp3),
	}
	options := &azopenai.GenerateSpeechFromTextOptions{}

	slog.DebugContext(ctx, "Generating speech from text", slog.String("input", *body.Input), slog.String("voice", string(*body.Voice)), slog.String("format", string(*body.ResponseFormat)))

	res, err := mgr.OpenAPI.GenerateSpeechFromText(ctx, body, options)
	if err != nil {
		return errors.Errorf("failed to generate speech: %w", err)
	}

	outFile, err := os.Create("speech.mp3")
	if err != nil {
		return errors.Errorf("failed to create output file: %w", err)
	}

	defer outFile.Close()
	slog.DebugContext(ctx, "Writing speech to file", slog.String("file", "speech.mp3"))

	_, err = io.Copy(outFile, res.Body)
	if err != nil {
		return errors.Errorf("failed to write to file: %w", err)
	}

	slog.InfoContext(ctx, "Speech file created", slog.String("file", "speech.mp3"))

	return nil
}
