package demo

import (
	"context"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/angelokurtis/go-talk/internal/errors"
)

type MP3Writer struct{}

func NewMP3Writer() *MP3Writer {
	return &MP3Writer{}
}

// Write takes an io.Reader containing MP3 data and writes it to a file with a timestamped name.
func (fw *MP3Writer) Write(ctx context.Context, data io.Reader) error {
	// Generate a filename with current timestamp.
	timestamp := time.Now().Format("20060102-150405")

	wd, err := os.Getwd()
	if err != nil {
		return errors.WithStack(err)
	}

	filename := filepath.Join(wd, timestamp+".mp3")
	slog.InfoContext(ctx, "Writing MP3 file", slog.String("filename", filename))

	// Create and open the file.
	file, err := os.Create(filename)
	if err != nil {
		return errors.WithStack(err)
	}

	defer file.Close()

	// Write the data to the file.
	if _, err = io.Copy(file, data); err != nil {
		return errors.WithStack(err)
	}

	slog.DebugContext(ctx, "MP3 file written successfully", slog.String("filename", filename))

	return nil
}
