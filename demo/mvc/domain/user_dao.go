package domain

import (
	"fmt"
	"github.com/golanshy/golang-microservices/demo/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Golan", LastName: "Shay", Email: "my_email@gmail.com"},
	}

	UserDau usersDaoInterface
)

func init() {
	UserDau = &userDao{}
}

type usersDaoInterface interface {
	GetUser(userId int64) (*User, *utils.ApplicationError)
}

type userDao struct {}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("We're accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("user %v not found", userId),
		StatusCode:  http.StatusNotFound,
		Code:    "not_found",
	}
}
