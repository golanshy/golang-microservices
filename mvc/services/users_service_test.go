package services

import (
	"github.com/golanshy/golang-microservices/mvc/domain"
	"github.com/golanshy/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	userDaoMock usersDaoMock
	getUserFunction func (userId int64)  (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDau = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}
func TestUserService_GetUserNotFoundInDB(t *testing.T) {
	getUserFunction = func(userId int64) (user *domain.User, applicationError *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message: "user 0 not found",
			StatusCode:  http.StatusNotFound,
			Code:    "not_found",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 not found", err.Message)
}

func TestUserService_GetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (user *domain.User, applicationError *utils.ApplicationError) {
		return &domain.User {
			Id: 123,
		}, nil
	}
	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
}
