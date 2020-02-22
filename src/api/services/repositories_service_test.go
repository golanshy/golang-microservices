package services

import (
	"github.com/golanshy/golang-microservices/src/api/clients/restClient"
	"github.com/golanshy/golang-microservices/src/api/domain/repositories"
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

func TestCreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
	assert.EqualValues(t, "Invalid repository name", err.Message())
}

func TestCreateRepoFailedErrorFromGithub(t *testing.T) {
	restClient.FlushMockUps()
	restClient.AddMockUp(restClient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Status:           "",
			StatusCode:       http.StatusUnauthorized,
			Proto:            "",
			ProtoMajor:       0,
			ProtoMinor:       0,
			Header:           nil,
			Body:             ioutil.NopCloser(strings.NewReader(`{"message": "Requires Authentication", "documentation_ur": "https://developer...."}`)),
			ContentLength:    0,
			TransferEncoding: nil,
			Close:            false,
			Uncompressed:     false,
			Trailer:          nil,
			Request:          nil,
			TLS:              nil,
		},
		Err: nil,
	})
	request := repositories.CreateRepoRequest{
		Name:        "Test repo",
		Description: "Repo Description",
	}
	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires Authentication", err.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	restClient.FlushMockUps()
	restClient.AddMockUp(restClient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			Status:           "",
			StatusCode:       http.StatusCreated,
			Proto:            "",
			ProtoMajor:       0,
			ProtoMinor:       0,
			Header:           nil,
			Body:             ioutil.NopCloser(strings.NewReader(`{"id": 123, "name" : "Test repo", "owner": {"login": "golanshy"}}`)),
			ContentLength:    0,
			TransferEncoding: nil,
			Close:            false,
			Uncompressed:     false,
			Trailer:          nil,
			Request:          nil,
			TLS:              nil,
		},
		Err: nil,
	})

	request := repositories.CreateRepoRequest{
		Name:        "Test repo",
		Description: "Repo Description",
	}
	result, err := RepositoryService.CreateRepo(request)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "Test repo", result.Name)
	assert.EqualValues(t, "golanshy", result.Owner)
}
