// +build integration

package presenter

import (
	"context"
	"net/http"
	"testing"

	"gotest.tools/assert"
)

// all of middleware should be tested

func TestUtf8Valid(t *testing.T) {
	username := "\xD0\x9D\xD0\xB8\xD0\xBA\xD0\xBE\xD0\xBB\xD0\xB0"
	handlerToTest := wrappedUtf8ValidatorHandler(username, func(w http.ResponseWriter, r *http.Request) {})

	testHandlerFunc(t, handlerToTest, "http://testing", http.StatusOK)
}

func TestUtf8Invalid(t *testing.T) {
	username := "\xD0\x9D\xD0\xB8\xD0\xBA\xD0\xBE\xD0\xBB\xD0"
	handlerToTest := wrappedUtf8ValidatorHandler(username, func(w http.ResponseWriter, r *http.Request) {})

	respBytes := testHandlerFunc(t, handlerToTest, "http://testing", http.StatusNotAcceptable)
	errStruct := unmarshallErrStruct(t, respBytes)

	expectedErrMsg := "username: " + ErrInvalidString.Error()
	assert.Assert(t, errStruct.Msg == expectedErrMsg)
}

func wrappedUtf8ValidatorHandler(username string, next http.HandlerFunc) http.HandlerFunc {
	envelope := func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			newCtx := context.WithValue(r.Context(), ctxKey("username"), username)
			h(w, r.WithContext(newCtx))
		}
	}

	return envelope(utf8Validator(next))
}
