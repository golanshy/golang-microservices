package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/golanshy/golang-microservices/src/api/domain/repositories"
	"github.com/golanshy/golang-microservices/src/api/services"
	"github.com/golanshy/golang-microservices/src/api/utils/errors"
	"net/http"
)

func CreateRepo(c *gin.Context) {
	//isPrivate := c.GetHeader("X-Private")
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	//clientId := c.Param("X-ClientID")
	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func CreateRepos(c *gin.Context) {
	var request []repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	//clientId := c.Param("X-ClientID")
	result, err := services.RepositoryService.CreateRepos(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(result.StatusCode, result)
}
