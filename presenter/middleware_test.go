package presenter

import (
	"net/http"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"gotest.tools/assert"
)

func TestI2StringNil(t *testing.T) {
	key := "route"
	_, status, err := interfaceToString(nil, key)
	assert.Assert(t, errors.Is(err, errMissingDependecy))
	assert.Assert(t, strings.Contains(err.Error(), key))
	assert.Assert(t, status == http.StatusInternalServerError)
}

func TestI2StringNotString(t *testing.T) {
	key := "route"
	inter := []string{"route"}
	_, status, err := interfaceToString(inter, key)
	assert.Assert(t, errors.Is(err, errMissingDependecy))
	assert.Assert(t, strings.Contains(err.Error(), key))
	assert.Assert(t, status == http.StatusInternalServerError)
}

func TestI2StringOK(t *testing.T) {
	key := "route"
	inter := "/hello/"
	s, _, err := interfaceToString(inter, key)
	assert.NilError(t, err)
	assert.Assert(t, s == inter)
}
