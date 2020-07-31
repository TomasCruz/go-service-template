// +build integration

package presenter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func testHandlerFunc(t *testing.T,
	handlerToTest http.HandlerFunc,
	target string,
	expectedStatus int) (respBytes []byte) {

	r := httptest.NewRecorder()
	req := httptest.NewRequest("GET", target, nil)
	handlerToTest(r, req)

	resp := r.Result()
	status := resp.StatusCode
	assert.Assert(t, status == expectedStatus, "wrong status code: got %v, expected %v", status, expectedStatus)

	var err error
	respBytes, err = ioutil.ReadAll(resp.Body)
	assert.NilError(t, err)

	return
}

func unmarshallMsgStruct(t *testing.T, respBytes []byte) MsgStruct {
	msg := MsgStruct{}
	err := json.Unmarshal(respBytes, &msg)
	assert.NilError(t, err)

	return msg
}

func unmarshallErrStruct(t *testing.T, respBytes []byte) ErrStruct {
	errStruct := ErrStruct{}
	err := json.Unmarshal(respBytes, &errStruct)
	assert.NilError(t, err)

	return errStruct
}
