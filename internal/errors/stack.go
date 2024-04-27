package errors

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/lmittmann/tint"
)

// Stack represents a call stack, a slice of function call program counter (PC) values
type Stack []uintptr

// callers function retrieves the call stack information
func callers() *Stack {
	s := make(Stack, 64)
	n := runtime.Callers(3, s)
	s = s[:n]

	return &s
}

// Format method on Stack type formats the call stack into a human-readable string
func (s *Stack) Format() string {
	if s == nil {
		return ""
	}

	// Create a new stackBuilder to accumulate formatted frame information
	builder := new(stackBuilder)
	frames := runtime.CallersFrames(*s)

	for {
		frame, more := frames.Next()

		// Check if the frame belongs to the runtime package (e.g., runtime.Callers)
		if strings.HasSuffix(filepath.Dir(frame.File), "/runtime") {
			return builder.Build()
		}

		// If it's not a runtime frame, add it to the builder
		builder.AddCallerFrame(frame)

		// Exit the loop if there are no more frames
		if !more {
			return builder.Build()
		}
	}
}

// stackBuilder is a helper struct to build the formatted stack trace string
type stackBuilder struct {
	builder strings.Builder
}

// AddCallerFrame method on stackBuilder adds a single frame information to the string
func (s *stackBuilder) AddCallerFrame(frame runtime.Frame) {
	if _, err := s.builder.WriteString(fmt.Sprintf(
		"\n%s\n\t%s:%d",
		frame.Function,
		frame.File,
		frame.Line,
	)); err != nil {
		slog.WarnContext(context.TODO(), "Error formatting stack trace", tint.Err(err))
	}
}

// Build method on stackBuilder returns the final formatted stack trace string
func (s *stackBuilder) Build() string {
	return s.builder.String()
}
