package errors

import (
	"errors"
	"fmt"
)

// New creates a new error with the provided message and captures the current call stack.
func New(msg string) error {
	return &wrapper{
		err:   fmt.Errorf(msg),
		stack: callers(),
	}
}

// WithStack wraps an existing error with a captured call stack.
// It attempts to reuse the existing stack trace if the error is already wrapped using this package (`errors.wrapper`).
func WithStack(err error) error {
	if err == nil {
		return nil
	}

	if w := new(wrapper); errors.As(err, &w) {
		return err
	}

	return &wrapper{
		err:   err,
		stack: callers(),
	}
}

// Errorf creates a new error using fmt.Errorf and attempts to capture the call stack from any argument that implements the Traceable interface.
// If no argument implements Traceable, it falls back to capturing the current call stack.
func Errorf(msg string, args ...any) error {
	e := fmt.Errorf(msg, args...)

	for _, arg := range args {
		if t, ok := arg.(Traceable); ok {
			return &wrapper{
				err:   e,
				stack: t.Stack(),
			}
		}
	}

	return &wrapper{
		err:   e,
		stack: callers(),
	}
}

// wrapper struct is used to wrap an underlying error and store its call stack information.
type wrapper struct {
	err   error
	stack *Stack
}

// Error implements the built-in error interface.
// It returns the string representation of the underlying error wrapped by this instance.
func (t *wrapper) Error() string {
	if t.err == nil {
		return ""
	}

	return t.err.Error()
}

// Cause retrieves the cause (original) error wrapped by this instance.
func (t *wrapper) Cause() error {
	return t.err
}

// Stack retrieves the call stack information associated with this error.
func (t *wrapper) Stack() *Stack {
	return t.stack
}

// String returns a formatted string representation of the error message and its call stack.
func (t *wrapper) String() string {
	return t.Stack().Format()
}
