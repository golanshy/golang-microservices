package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/golanshy/golang-microservices/src/api/clients/restClient"
	"github.com/golanshy/golang-microservices/src/api/domain/repositories"
	"github.com/golanshy/golang-microservices/src/api/utils/errors"
	"github.com/golanshy/golang-microservices/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	restClient.StartMockUps()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidJsonRequest(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))
	response := httptest.NewRecorder()
	c := test_utils.GetMockContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	fmt.Println(response.Body.String())

	apiErr, err := errors.NewAPiErrorFoBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "Invalid json body", apiErr.Message())
}

func TestCreateRepoErrorFomGithub(t *testing.T) {
	restClient.FlushMockUps()
	restClient.AddMockUp(restClient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires Authentication", "documentation_url": "https:/developer..."}`)),
		},
	})

	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockContext(request, response)

	CreateRepo(c)

	assert.EqualValues(t, http.StatusUnauthorized, response.Code)
	fmt.Println(response.Body.String())

	apiErr, err := errors.NewAPiErrorFoBytes(response.Body.Bytes())

	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusUnauthorized, apiErr.Status())
	assert.EqualValues(t, "Requires Authentication", apiErr.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	restClient.FlushMockUps()
	restClient.AddMockUp(restClient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "Hello-World", "full_name": "golanshy/Hello-World"}`)),
		},
	})

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

