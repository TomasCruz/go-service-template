package presenter

import (
	"net/http"
)

// DummyMiddleware does nothing
func DummyMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}
