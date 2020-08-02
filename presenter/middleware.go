package presenter

import (
	"context"
	"net/http"
	"unicode/utf8"

	"github.com/TomasCruz/go-service-template/callstack"
	"github.com/pkg/errors"
)

type ctxKey string

func routeStorer(route string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newCtx := context.WithValue(r.Context(), ctxKey("route"), route)
		h(w, r.WithContext(newCtx))
	}
}

func userNameExtractor(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyString := "route"
		iRoute := r.Context().Value(ctxKey(keyString))
		route, status, err := interfaceToString(iRoute, keyString)
		if err != nil {
			callstack.PrintErrStack(err)
			errorResponse(w, err, status)
			return
		}

		length := len(route)
		usernameString := r.URL.Path[length:]
		if len(usernameString) > 0 && usernameString[len(usernameString)-1] == '/' {
			usernameString = usernameString[:len(usernameString)-1]
		}

		newCtx := context.WithValue(r.Context(), ctxKey("username"), usernameString)
		h(w, r.WithContext(newCtx))
	}
}

func utf8Validator(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyString := "username"
		iUsername := r.Context().Value(ctxKey(keyString))
		username, status, err := interfaceToString(iUsername, keyString)
		if err != nil {
			callstack.PrintErrStack(err)
			errorResponse(w, err, status)
			return
		}

		if !utf8.ValidString(username) {
			status := http.StatusNotAcceptable
			err := errors.Wrap(ErrInvalidString, keyString)
			errorResponse(w, err, status)
			return
		}

		h(w, r)
	}
}

func interfaceToString(i interface{}, key string) (s string, status int, err error) {
	s, ok := i.(string)
	if !ok {
		status = http.StatusInternalServerError
		err = errors.Wrap(errMissingDependecy, key)
	}

	return
}
