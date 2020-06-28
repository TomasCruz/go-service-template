// +build end2end

package main

import (
	"net/http"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestHealth(t *testing.T) {
	timeout := time.Duration(1000 * time.Millisecond)
	client := http.Client{Timeout: timeout}

	route := "http://localhost:4337/health/"
	req, err := http.NewRequest("GET", route, nil)
	assert.NilError(t, err)

	resp, err := client.Do(req)
	assert.NilError(t, err)
	defer resp.Body.Close()

	assert.Assert(t, 204 == resp.StatusCode)
}
