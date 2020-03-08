package repositories

import (
	"github.com/golanshy/golang-microservices/go-common/utils/errors"
	"strings"
)

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateRepoRequest) Validate() errors.APiError {
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return errors.NewBadRequestError("Invalid repository name")
	}
	return nil
}

type CreateRepoResponse struct {
	Id    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type CreateReposResponse struct {
	StatusCode int                      `json:"status"`
	Results    []CreateRepositoriesResult `json:"results"`
}

type CreateRepositoriesResult struct {
	Response *CreateRepoResponse `json:"owner"`
	Error    errors.APiError     `json:"error"`
}
