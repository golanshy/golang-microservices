package providers

import (
	"errors"
	"github.com/golanshy/golang-microservices/go-common/src/api/clients/restClient"
	"github.com/golanshy/golang-microservices/oauth-api/domain/repositories"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	restClient.StartMockUps()
	os.Exit(m.Run())
}

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "content-type", contentType)
	assert.EqualValues(t, "application/x-www-form-urlencoded", contentTypeApplication)
	assert.EqualValues(t, "https://applylogic.auth0.com/oauth/token", urlCreateSessionToken)
	assert.EqualValues(t, "authorization", headerAuthorization)
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "Bearer abc123", header)
}

func TestCreateSessionTokenErrorRestClient(t *testing.T) {
	restClient.FlushMockUps()
	restClient.AddMockUp(restClient.Mock{
		Url:        "https://applylogic.auth0.com/oauth/token",
		HttpMethod: http.MethodPost,
		Response:   nil,
		Err:        errors.New("invalid restClient response"),
	})
	response, err := CreateSessionToken(repositories.CreateSessionTokenRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid restClient response", err.ErrorDescription)
}

func TestCreateSessionTokenInvalidResponseBody(t *testing.T) {
	restClient.FlushMockUps()
	invalidCloser, _ := os.Open("-asf3")
	restClient.AddMockUp(restClient.Mock{
		Url:        "https://applylogic.auth0.com/oauth/token",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       invalidCloser,
		},
	})
	response, err := CreateSessionToken(repositories.CreateSessionTokenRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid response body", err.ErrorDescription)
}

func TestCreateSessionTokenInvalidErrorInterface(t *testing.T) {
	restClient.FlushMockUps()

	restClient.AddMockUp(restClient.Mock{
		Url:        "https://applylogic.auth0.com/oauth/token",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       ioutil.NopCloser(strings.NewReader(`{"error":"access_denied","error_description":"Unauthorized"}`)),
		},
	})
	response, err := CreateSessionToken(repositories.CreateSessionTokenRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid json response body", err.Error)
}

func TestCreateSessionTokenUnauthorized(t *testing.T) {
	restClient.FlushMockUps()

	restClient.AddMockUp(restClient.Mock{
		Url:        "https://applylogic.auth0.com/oauth/token",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"error":"access_denied","error_description":"Requires authentication"}`)),
		},
	})
	response, err := CreateSessionToken(repositories.CreateSessionTokenRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "access_denied", err.Error)
	assert.EqualValues(t, "Requires authentication", err.ErrorDescription)
}

func TestCreateSessionTokenSuccessInvalidResponse(t *testing.T) {
	restClient.FlushMockUps()

	restClient.AddMockUp(restClient.Mock{
		Url:        "https://applylogic.auth0.com/oauth/token",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": "123"}`)),
		},
	})
	response, err := CreateSessionToken(repositories.CreateSessionTokenRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "error when trying to unmarshal create session token response", err.ErrorDescription)
}

func TestCreateSuccessNoError(t *testing.T) {
	restClient.FlushMockUps()

	restClient.AddMockUp(restClient.Mock{
		Url:        "https://applylogic.auth0.com/oauth/token",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"access_token": "access_token_test", "token_type": "token_type_test"}`)),
		},
	})
	response, err := CreateSessionToken(repositories.CreateSessionTokenRequest{})
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, "access_token_test", response.AccessToken)
	assert.EqualValues(t, "token_type_test", response.TokenType)
}