// +build integration

package presenter

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestHealth(t *testing.T) {
	handlerToTest := HealthHandler

	r := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://testing", nil)
	handlerToTest(r, req)

	status := r.Code
	expectedStatus := http.StatusNoContent
	assert.Assert(t, status == expectedStatus, "wrong status code: got %v, expected %v", status, expectedStatus)
}

func TestHello(t *testing.T) {
	handlerToTest := routeStorer(
		"/hello/", userNameExtractor(
			utf8Validator(
				hello)))

	respBytes := testHandlerFunc(t, handlerToTest, "http://testing/hello/JohnDoe", http.StatusOK)
	msg := unmarshallMsgStruct(t, respBytes)

	expectedMsg := "hello JohnDoe"
	assert.Assert(t, msg.Msg == expectedMsg, "unexpected body: got %v, expected %v", msg.Msg, expectedMsg)
}

func TestHelloWorld(t *testing.T) {
	handlerToTest := routeStorer(
		"/hello/", userNameExtractor(
			utf8Validator(
				hello)))

	respBytes := testHandlerFunc(t, handlerToTest, "http://testing/hello/", http.StatusOK)
	msg := unmarshallMsgStruct(t, respBytes)

	expectedMsg := "hello world"
	assert.Assert(t, msg.Msg == expectedMsg, "unexpected body: got %v, expected %v", msg.Msg, expectedMsg)
}
