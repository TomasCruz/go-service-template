package presenter

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/TomasCruz/go-service-template/callstack"
	"github.com/TomasCruz/go-service-template/service"
)

// HealthHandler displays health status of the service.
// Status NoContent (204) is returned if service is working as expected.
// Status InternalServerError (500) is returned in case of general errors.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if err := service.Health(); err != nil {
		callstack.LogErrStack(err)
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
	length := len(rts.HelloRoute)
	usernameString := r.URL.Path[length:]
	if len(usernameString) > 0 && usernameString[len(usernameString)-1] == '/' {
		usernameString = usernameString[:len(usernameString)-1]
	}

	msgString, err := service.Hello(usernameString)
	if err != nil {
		status := http.StatusInternalServerError
		callstack.LogErrStack(err)

		if errors.Is(err, service.ErrHello) {
			status = http.StatusNotAcceptable
			err = service.ErrHello
		} else if errors.Is(err, service.ErrInvalidString) {
			status = http.StatusUnprocessableEntity
			err = service.ErrInvalidString
		} else {
			err = errInternalServerError
		}

		errorResponse(w, err, status)
		return
	}

	msg := MsgStruct{Msg: msgString}
	data, err := json.Marshal(&msg)
	if err != nil {
		callstack.LogErrStack(err)
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
