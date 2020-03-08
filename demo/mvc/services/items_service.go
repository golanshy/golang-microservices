package services

import (
	"github.com/golanshy/golang-microservices/demo/mvc/domain"
	"github.com/golanshy/golang-microservices/demo/mvc/utils"
	"net/http"
)
type itemService struct {}

var (ItemService itemService)

func (u *itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message: "implement me",
		StatusCode: http.StatusInternalServerError,
		Code: "implement_me",
	}
}
