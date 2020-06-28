// +build end2end

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/TomasCruz/go-service-template/service"

	"github.com/TomasCruz/go-service-template/presenter"

	"gotest.tools/assert"
)

func TestHello(t *testing.T) {
	timeout := time.Duration(1000 * time.Millisecond)
	client := http.Client{Timeout: timeout}

	route := "http://localhost:4337/hello/Toma"
	req, err := http.NewRequest("GET", route, nil)
	assert.NilError(t, err)

	resp, err := client.Do(req)
	assert.NilError(t, err)
	defer resp.Body.Close()

	assert.Assert(t, 200 == resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.NilError(t, err)

	var msgStruct presenter.MsgStruct
	err = json.Unmarshal(bodyBytes, &msgStruct)
	assert.NilError(t, err)

	assert.Assert(t, "hello Toma" == msgStruct.Msg)
}

func TestInvalidString(t *testing.T) {
	timeout := time.Duration(1000 * time.Millisecond)
	client := http.Client{Timeout: timeout}

	route := "http://localhost:4337/hello/\xD0\x9D\xD0\xB8\xD0\xBA\xD0\xBE\xD0\xBB\xD0"
	req, err := http.NewRequest("GET", route, nil)
	assert.NilError(t, err)

	resp, err := client.Do(req)
	assert.NilError(t, err)
	defer resp.Body.Close()

	assert.Assert(t, 500 == resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.NilError(t, err)

	var errStruct presenter.ErrStruct
	err = json.Unmarshal(bodyBytes, &errStruct)
	assert.NilError(t, err)

	assert.Assert(t, service.ErrInvalidString.Error() == errStruct.Msg)
}

func TestStar(t *testing.T) {
	timeout := time.Duration(1000 * time.Millisecond)
	client := http.Client{Timeout: timeout}

	route := "http://localhost:4337/hello/*"
	req, err := http.NewRequest("GET", route, nil)
	assert.NilError(t, err)

	resp, err := client.Do(req)
	assert.NilError(t, err)
	defer resp.Body.Close()

	assert.Assert(t, 500 == resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.NilError(t, err)

	var errStruct presenter.ErrStruct
	err = json.Unmarshal(bodyBytes, &errStruct)
	assert.NilError(t, err)

	assert.Assert(t, service.ErrHello.Error() == errStruct.Msg)
}
