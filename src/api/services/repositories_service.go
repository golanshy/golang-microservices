package services

import (
	"github.com/golanshy/golang-microservices/src/api/config"
	"github.com/golanshy/golang-microservices/src/api/domain/github"
	"github.com/golanshy/golang-microservices/src/api/domain/repositories"
	"github.com/golanshy/golang-microservices/src/api/providers/github_provider"
	"github.com/golanshy/golang-microservices/src/api/utils/errors"
	"net/http"
	"sync"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APiError)
	CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APiError) {

	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreatRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}
	//option_a.Info("about to send request to external api",
	//	fmt.Sprintf("clientId:%s", cliendId),
	//	"status:pending")
	//option_b.Info("about to send request to external api",
	//	option_b.Field("clientId", cliendId),
	//	option_b.Field("status", "pending"),
	//	option_b.Field("authenticated", cliendId != "")

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		//option_a.Error("response obtained from external api",
		//	err,
		//	fmt.Sprintf("clientId:%s", cliendId),
		//	"status:error")
		//option_b.Error("response obtained from external api",
		//	err,
		//	option_b.Field("clientId", cliendId),
		//	option_b.Field("status", "error"),
		//	option_b.Field("authenticated", cliendId != ""))

		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}
	//option_a.Info("response obtained from external api",
	//	fmt.Sprintf("clientId:%s", cliendId),
	//	"status:success")
	//option_b.Info("response obtained from external api",
	//	option_b.Field("clientId", cliendId),
	//	option_b.Field("status", "success"),
	//	option_b.Field("authenticated", cliendId != "")

	result := &repositories.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return result, nil
}

func (s *repoService) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.APiError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, input, output)

	for _, current := range request {
		wg.Add(1)
		go s.createRepoConcurrent(current, input)
	}

	wg.Wait()
	close(input)
	results := <-output

	successCreations := 0
	for _, current := range results.Results {
		if current.Response != nil {
			successCreations++
		}
	}
	if successCreations == 0 {
		results.StatusCode = results.Results[0].Error.Status()
	} else if successCreations == len(request) {
		results.StatusCode = http.StatusCreated
	} else {
		results.StatusCode = http.StatusPartialContent
	}
	return results, nil
}

func (s *repoService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse
	for incomingEvent := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingEvent.Response,
			Error:    incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}
	output <- results
}

func (s *repoService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	result, err := s.CreateRepo(input)
	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	output <- repositories.CreateRepositoriesResult{Response: result}
}
