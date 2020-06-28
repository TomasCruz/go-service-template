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

func TestInvalid(t *testing.T) {
	_, err := Hello("\xD0\x9D\xD0\xB8\xD0\xBA\xD0\xBE\xD0\xBB\xD0")
	assert.Assert(t, errors.Is(err, ErrInvalidString))
}
