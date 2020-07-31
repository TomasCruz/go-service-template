package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/TomasCruz/go-service-template/callstack"
	"github.com/TomasCruz/go-service-template/service"
)

// HealthHandler displays health status of the service.
// Status NoContent (204) is returned if service is working as expected.
// Status InternalServerError (500) is returned in case of general errors.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if err := service.Health(); err != nil {
		callstack.PrintErrStack(err)
		errorResponse(w, nil, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// HelloHandler says hello
// Status OK (200) is returned for successfuly saying hello.
// Status NotAcceptable (406) is returned for unacceptable input.
// Status UnprocessableEntity (422) is returned for invalid input.
// Status InternalServerError (500) is returned in case of general errors.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	retHandler := routeStorer(
		rts.HelloRoute, userNameExtractor(
			utf8Validator(
				hello)))

	retHandler(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	iUsername := r.Context().Value(ctxKey("username"))
	username := iUsername.(string)

	msgString, err := service.Hello(username)
	if err != nil {
		callstack.PrintErrStack(err)

		status := http.StatusInternalServerError
		if errors.Is(err, service.ErrHello) {
			status = http.StatusNotAcceptable
		} else {
			err = errInternalServerError
		}

		errorResponse(w, err, status)
		return
	}

	msg := MsgStruct{Msg: msgString}
	data, err := json.Marshal(&msg)
	if err != nil {
		callstack.PrintErrStack(err)
		errorResponse(w, errInternalServerError, http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func errorResponse(w http.ResponseWriter, err error, errCode int) {
	if err != nil {
		// if there is also an error with JSON marshalling, return 500
		data, err := json.Marshal(ErrStruct{Msg: err.Error()})
		if err != nil {
			errCode = http.StatusInternalServerError
		}

		w.WriteHeader(errCode)
		w.Write(data)
		return
	}

	w.WriteHeader(errCode)
}
