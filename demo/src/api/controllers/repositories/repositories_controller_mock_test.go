package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/golanshy/golang-microservices/demo/src/api/domain/repositories"
	"github.com/golanshy/golang-microservices/demo/src/api/services"
	"github.com/golanshy/golang-microservices/go-common/utils/errors"
	"github.com/golanshy/golang-microservices/go-common/utils/test_utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	funcCreateRepo  func(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APiError)
	funcCreateRepos func(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APiError)
)

type repoServiceMock struct {
}

func (s *repoServiceMock) CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APiError) {
	return funcCreateRepo(request)
}
func (s *repoServiceMock) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APiError) {
	return funcCreateRepos(request)
}

func TestCreateRepoNoErrorMockingTheEntireService(t *testing.T) {
	services.RepositoryService = &repoServiceMock{}

	funcCreateRepo = func(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APiError) {
		return &repositories.CreateRepoResponse{
			Id: 123,
			Name: "Hello-World",
			Owner: "",
		}, nil
	}

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)
	fmt.Println(response.Body.String())

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 123, result.Id)
	assert.EqualValues(t, "Hello-World", result.Name)
	assert.EqualValues(t, "", result.Owner)
}
