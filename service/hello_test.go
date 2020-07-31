package service

import (
	"errors"
	"testing"

	"gotest.tools/assert"
)

func TestEmpty(t *testing.T) {
	msg, err := Hello("")
	assert.NilError(t, err)

	assert.Assert(t, "hello world" == msg)
}

func TestStar(t *testing.T) {
	_, err := Hello("*")
	assert.Assert(t, errors.Is(err, ErrHello))
}

func TestToma(t *testing.T) {
	msg, err := Hello("Toma")
	assert.NilError(t, err)

	assert.Assert(t, "hello Toma" == msg)
}
