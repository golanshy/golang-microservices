package services

import (
	"github.com/golanshy/golang-microservices/demo/src/api/domain/repositories"
	"github.com/golanshy/golang-microservices/go-common/src/api/clients/restClient"
	"github.com/golanshy/golang-microservices/go-common/utils/errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
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
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires Authentication", "documentation_ur": "https://developer...."}`)),
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
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name" : "Test repo", "owner": {"login": "golanshy"}}`)),
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

func TestCreateRepoConcurrentInvalidRequest(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	output := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.createRepoConcurrent(request, output)

	result := <-output

	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
	assert.EqualValues(t, "Invalid repository name", result.Error.Message())
}

func TestCreateRepoConcurrentErrorFromGithub(t *testing.T) {

	restClient.FlushMockUps()
	restClient.AddMockUp(restClient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires Authentication", "documentation_ur": "https://developer...."}`)),
		},
		Err: nil,
	})

	request := repositories.CreateRepoRequest{
		Name:        "Test repo",
		Description: "Repo Description",
	}

	output := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.createRepoConcurrent(request, output)

	result := <-output

	assert.NotNil(t, result)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusUnauthorized, result.Error.Status())
	assert.EqualValues(t, "Requires Authentication", result.Error.Message())
}

func TestCreateRepoConcurrentNoError(t *testing.T) {
	restClient.FlushMockUps()
	restClient.AddMockUp(restClient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name" : "Test repo", "owner": {"login": "golanshy"}}`)),
		},
		Err: nil,
	})

	request := repositories.CreateRepoRequest{
		Name:        "Test repo",
		Description: "Repo Description",
	}

	output := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.createRepoConcurrent(request, output)

	result := <-output

	assert.NotNil(t, result)
	assert.NotNil(t, result.Response)
	assert.Nil(t, result.Error)
	assert.EqualValues(t, 123, result.Response.Id)
	assert.EqualValues(t, "Test repo", result.Response.Name)
	assert.EqualValues(t, "golanshy", result.Response.Owner)
}

func TestHandleResults(t *testing.T) {

	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)

	var wg sync.WaitGroup
	service := repoService{}
	go service.handleRepoResults(&wg, input, output)

	wg.Add(1)
	go func() {
		input <- repositories.CreateRepositoriesResult{
			Error: errors.NewBadRequestError("Invalid repository name"),
		}
	}()
	wg.Add(1)
	go func() {
		input <- repositories.CreateRepositoriesResult{
			Error: errors.NewBadRequestError("Invalid repository name"),
		}
	}()

	wg.Wait()
	close(input)

	result := <-output
	assert.NotNil(t, result)
	assert.EqualValues(t, 0, result.StatusCode)
	assert.EqualValues(t, 2, len(result.Results))

	assert.NotNil(t, result.Results[0].Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, "Invalid repository name", result.Results[0].Error.Message())
}

func TestCreateReposInvalidRequests(t *testing.T) {
	requests := []repositories.CreateRepoRequest{
		{},
		{Name: "  "},
	}

	result, err := RepositoryService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.Results)
	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)

	assert.EqualValues(t, 2, len(result.Results))

	assert.Nil(t, result.Results[0].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
	assert.EqualValues(t, "Invalid repository name", result.Results[0].Error.Message())

	assert.Nil(t, result.Results[1].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[1].Error.Status())
	assert.EqualValues(t, "Invalid repository name", result.Results[1].Error.Message())
}

func TestCreateReposOneSuccessOneFailed(t *testing.T) {

	restClient.FlushMockUps()
	restClient.AddMockUp(restClient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name" : "testing", "owner": {"login": "golanshy"}}`)),
		},
		Err: nil,
	})

	requests := []repositories.CreateRepoRequest{
		{},
		{Name: "testing"},
	}

	result, err := RepositoryService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.Results)
	assert.EqualValues(t, http.StatusPartialContent, result.StatusCode)

	assert.EqualValues(t, 2, len(result.Results))

	for _, result := range result.Results {
		if result.Error != nil {
			assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
			assert.EqualValues(t, "Invalid repository name", result.Error.Message())
			continue
		}

		assert.EqualValues(t, 123, result.Response.Id)
		assert.EqualValues(t, "testing", result.Response.Name)
		assert.EqualValues(t, "golanshy", result.Response.Owner)
	}
}

func TestCreateReposAllSuccess(t *testing.T) {

	restClient.FlushMockUps()
	restClient.AddMockUp(restClient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name" : "testing", "owner": {"login": "golanshy"}}`)),
		},
		Err: nil,
	})

	requests := []repositories.CreateRepoRequest{
		{Name: "Test repo"},
		{Name: "Test repo"},
	}

	result, err := RepositoryService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, result.Results)
	assert.EqualValues(t, http.StatusCreated, result.StatusCode)

	assert.EqualValues(t, 2, len(result.Results))

	assert.Nil(t, result.Results[0].Error)
	assert.EqualValues(t, 123, result.Results[0].Response.Id)
	assert.EqualValues(t, "testing", result.Results[0].Response.Name)
	assert.EqualValues(t, "golanshy", result.Results[0].Response.Owner)

	assert.Nil(t, result.Results[1].Error)
	assert.EqualValues(t, 123, result.Results[1].Response.Id)
	assert.EqualValues(t, "testing", result.Results[1].Response.Name)
	assert.EqualValues(t, "golanshy", result.Results[1].Response.Owner)
}
