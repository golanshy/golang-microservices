package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGithubError(t *testing.T) {
	// initialization
	githubError := GithubError{
		Resource: "Repository",
		Code:     "custom",
		Field:    "name",
		Message:  "name already exists on this account",
	}

	errorResponse := GithubErrorReposnse{
		StatusCode:       422,
		Message:          "Repository creation failed.",
		DocumentationUrl: "https://developer.github.com/v3/repos/#create",
		Errors: []GithubError{githubError},
	}

	// execution
	// Marshal takes an input interface and attempts to create a valid json string
	bytes, err := json.Marshal(errorResponse)

	// validation
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	fmt.Println(string(bytes))

	var target GithubErrorReposnse
	// Unmarshal takes an input byte array and a pointer that were trying to fill using this json
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.NotNil(t, target)
}
