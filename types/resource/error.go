package resource

import "fmt"

// Error wraps an error with a Resource
type Error struct {
	Err error
	*Resource
}

// Unwrap yields the underlying error for a Error.
func (e *Error) Unwrap() error { return e.Err }

func (e *Error) Error() string {
	return fmt.Sprintf("%v with %v", e.Err, e.Resource)
}
