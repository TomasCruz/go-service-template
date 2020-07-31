package service

import (
	"fmt"
	"sync"
)

var (
	// ErrHello is a message to be displayed to user wrapped in error
	ErrHello error

	wg sync.WaitGroup
)

func init() {
	// sentinels, no need to record it's call stacks
	ErrHello = fmt.Errorf("Saying hello failed")
}
