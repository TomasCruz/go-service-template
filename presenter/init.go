package presenter

import (
	"fmt"
)

var (
	errInternalServerError error

	// ErrInvalidString is a message to be displayed to user wrapped in error
	ErrInvalidString    error
	errMissingDependecy error
)

func init() {
	// sentinels, no need to record it's call stacks
	errInternalServerError = fmt.Errorf("Internal Server Error")
	ErrInvalidString = fmt.Errorf("invalid UTF-8 string")
	errMissingDependecy = fmt.Errorf("missing dependecy")
}
