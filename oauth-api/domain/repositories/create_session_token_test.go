package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSessionTokenRequestAsJson(t *testing.T) {

	// initialization
	request := CreateSessionTokenRequest{
		ClientId:        "client_id_test",
		ClientSecret: "client_secret_test",
	}

	// execution
	// Marshal takes an input interface and attempts to create a valid json string
	bytes, err := json.Marshal(request)

	// validation
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	fmt.Println(string(bytes))

	var target CreateSessionTokenRequest
	// Unmarshal takes an input byte array and a pointer that were trying to fill using this json
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.NotNil(t, target)
	assert.EqualValues(t,
		`{"client_id":"client_id_test","client_secret":"client_secret_test"}`,
		string(bytes))

	assert.EqualValues(t, target.ClientId, request.ClientId)
	assert.EqualValues(t, target.ClientSecret, request.ClientSecret)
}

func TestCreateSessionTokenResponseAsJson(t *testing.T) {

	//initialization
	response := CreateSessionTokenResponse{
		AccessToken:        "access_token_test",
		TokenType:     "token_type_test",
	}

	// execution
	// Marshal takes an input interface and attempts to create a valid json string
	bytes, err := json.Marshal(response)

	// validation
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	fmt.Println(string(bytes))

	var target CreateSessionTokenResponse
	// Unmarshal takes an input byte array and a pointer that were trying to fill using this json
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.NotNil(t, target)
	assert.EqualValues(t,
		`{"access_token":"access_token_test","token_type":"token_type_test"}`,
		string(bytes))

	assert.EqualValues(t, target.AccessToken, response.AccessToken)
	assert.EqualValues(t, target.TokenType, response.TokenType)
}