package domain

import (
	"fmt"
	"github.com/golanshy/golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Golan", LastName: "Shay", Email: "my_email@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v not found", userId),
		Status:  http.StatusNotFound,
		Code:    "not_found",
	}
}
