package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/golanshy/golang-microservices/src/api/clients/restClient"
	"github.com/golanshy/golang-microservices/src/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}
func CreateRepo(accessToken string, request github.CreatRepoRequest) (*github.CreatRepoResponse, *github.GithubErrorReposnse) {
	header := getAuthorizationHeader(accessToken)
	headers := http.Header{}
	headers.Set(headerAuthorization, header)

	response, err := restClient.Post(urlCreateRepo, request, headers)
	fmt.Println(response)
	if err != nil {
		log.Printf("error when trying to create new repo in github: %s", err.Error())
		return nil, &github.GithubErrorReposnse{
			StatusCode:       http.StatusInternalServerError,
			Message:          err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorReposnse{
			StatusCode:       http.StatusInternalServerError,
			Message:          "invalid response body",
		}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorReposnse
		if err:= json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorReposnse{
				StatusCode:       http.StatusInternalServerError,
				Message:          "invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreatRepoResponse
	if err:= json.Unmarshal(bytes, &result); err != nil {
		log.Printf("error when trying to unmarshall create repo successful response %s", err.Error())
		return nil, &github.GithubErrorReposnse{
			StatusCode:       http.StatusInternalServerError,
			Message:          "error when trying to unmarshal github create repo response",
		}
	}
	return &result, nil
}
