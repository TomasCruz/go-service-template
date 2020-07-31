package presenter

import (
	"context"
	"net/http"
	"unicode/utf8"

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
		ctx := r.Context()
		iRoute := ctx.Value(ctxKey("route"))
		route, ok := interfaceToString(iRoute, w)
		if !ok {
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
		iUsername := r.Context().Value(ctxKey("username"))
		username, ok := interfaceToString(iUsername, w)
		if !ok {
			return
		}

		if !utf8.ValidString(username) {
			status := http.StatusNotAcceptable
			err := errors.Wrap(ErrInvalidString, "username")
			errorResponse(w, err, status)
			return
		}

		h(w, r)
	}
}

func interfaceToString(i interface{}, w http.ResponseWriter) (iString string, ok bool) {
	iString, ok = i.(string)
	if !ok {
		status := http.StatusInternalServerError
		err := errors.Wrap(errMissingDependecy, "route")
		errorResponse(w, err, status)
		return
	}

	return
}
