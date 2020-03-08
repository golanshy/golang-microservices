package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/golanshy/golang-microservices/go-common/utils/errors"
	"github.com/golanshy/golang-microservices/oauth-api/domain/repositories"
	"github.com/golanshy/golang-microservices/oauth-api/services"
	"net/http"
)

func CreateSessionToken(c *gin.Context) {
	//isPrivate := c.GetHeader("X-Private")
	var request repositories.CreateSessionTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	//clientId := c.Param("X-ClientID")
	result, err := services.RepositoryService.CreateSessionToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}
