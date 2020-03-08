package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T) {

	// initialization
	request := CreatRepoRequest{
		Name:        "golang_introduction",
		Description: "a golang introduction repository",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   false,
		HasProjects: true,
		HasWiki:     false,
	}

	// execution
	// Marshal takes an input interface and attempts to create a valid json string
	bytes, err := json.Marshal(request)

	// validation
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	fmt.Println(string(bytes))

	var target CreatRepoRequest
	// Unmarshal takes an input byte array and a pointer that were trying to fill using this json
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.NotNil(t, target)

	//assert.EqualValues(t,
	//	`{"name":"golang_introduction","description":"a golang introduction repository","homepage":"https://github.com","private":true,"has_issues":false,"has_projects":true,"has_wiki":false}`,
	//	string(bytes))

	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.Description, request.Description)
	assert.EqualValues(t, target.Homepage, request.Homepage)
	assert.EqualValues(t, target.Private, request.Private)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
	assert.EqualValues(t, target.HasProjects, request.HasProjects)
	assert.EqualValues(t, target.HasWiki, request.HasWiki)
}

func TestCreateRepoResponseAsJson(t *testing.T) {

	// initialization

	permissions := RepoPermissions{
		IsAdmin: true,
		HasPush: true,
		HasPull: true,
	}

	owner := RepoOwner{
		Id:        123,
		Login:     "golanshy",
		AvatarUrl: "https://avatars0.githubusercontent.com/u/1773923?v=4",
		Url:       "https://api.github.com/users/golanshy",
		HtmlUrl:   "https://github.com/golanshy",
	}

	response := CreatRepoResponse{
		Id:        240775830,
		Name:     "golang_introduction",
		FullName: "golanshy/golang_introduction",
		Owner:       owner,
		Permissions:   permissions,
	}

	// execution
	// Marshal takes an input interface and attempts to create a valid json string
	bytes, err := json.Marshal(response)

	// validation
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	fmt.Println(string(bytes))

	var target CreatRepoResponse
	// Unmarshal takes an input byte array and a pointer that were trying to fill using this json
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.NotNil(t, target)
}
