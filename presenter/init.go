package presenter

import (
	"fmt"
)

var errInternalServerError error

func init() {
	// sentinel, no need to record it's call stacks
	errInternalServerError = fmt.Errorf("Internal Server Error")
}
