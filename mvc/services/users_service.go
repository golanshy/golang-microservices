package services

import (
	"github.com/golanshy/golang-microservices/mvc/domain"
	"github.com/golanshy/golang-microservices/mvc/utils"
)

type userService struct {}

var (UserService userService)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDau.GetUser(userId)
}
