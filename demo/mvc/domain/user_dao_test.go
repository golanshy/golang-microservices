package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUseFound(t *testing.T) {

	// Initialization/

	// Execution
	user, err := GetUser(0)

	// Validation

	if user != nil {
		t.Error("We were not expecting a user with user_id := 0")
	}

	if err == nil {
		t.Error("We were expecting an error when user_id := 0")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("We were expecting http.StatusNotFound")
	}

	assert.Nil(t, user, "We were not expecting a user with user_id := 0")
	assert.NotNil(t, err, "We were expecting an error when user_id := 0")
	assert.EqualValues(t, "user 0 not found", err.Message, )
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)

}

func TestGetUserFound(t *testing.T) {

	// Initialization/

	// Execution
	user, err := GetUser(123)

	// Validation

	if user == nil {
		t.Error("We were not expecting a user with user_id := 123")
	}

	if err != nil {
		t.Error("We were expecting an error when user_id := 123")
	}

	assert.NotNil(t, user, "We were not expecting a user with user_id := 123")
	assert.Nil(t, err, "We were expecting an error when user_id := 123")
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Golan", user.FirstName)
	assert.EqualValues(t, "Shay", user.LastName)
	assert.EqualValues(t, "my_email@gmail.com", user.Email)
}
