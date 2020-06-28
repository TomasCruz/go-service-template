package service

import (
	"fmt"
	"sync"
)

var (
	// ErrHello is a message to be displayed to user wrapped in error
	ErrHello error

	// ErrInvalidString is a message to be displayed to user wrapped in error
	ErrInvalidString error

	wg sync.WaitGroup
)

func init() {
	// exported errors, being sentinels, don't need to record their call stacks
	ErrHello = fmt.Errorf("Saying hello failed")
	ErrInvalidString = fmt.Errorf("invalid UTF-8 string")
}
