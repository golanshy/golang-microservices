package test_utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMockContext(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "http://localhost:123/something", nil)
	assert.Nil(t, err)
	assert.NotNil(t, request)

	request.Header = http.Header{"X-Mock": {"true"},}

	response := httptest.NewRecorder()


	c := GetMockContext(request, response)

	assert.EqualValues(t, http.MethodGet, c.Request.Method)
	assert.EqualValues(t, "123", c.Request.URL.Port())
	assert.EqualValues(t, "/something", c.Request.URL.Path)
	assert.EqualValues(t, "http", c.Request.URL.Scheme)
	assert.EqualValues(t, 1, len(c.Request.Header))
	assert.EqualValues(t, "true", c.GetHeader("X-Mock"))

}
