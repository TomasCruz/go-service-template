package service

import (
	"fmt"

	"github.com/pkg/errors"
)

// Hello returns a greeting to username and an error
func Hello(username string) (msg string, err error) {
	wg.Add(1)
	defer wg.Done()

	if username == "*" {
		err = errors.Wrap(ErrHello, "Can't say hi to everybody")
		return
	}

	if username == "" {
		username = "world"
	}

	msg = fmt.Sprintf("hello %s", username)
	return
}
