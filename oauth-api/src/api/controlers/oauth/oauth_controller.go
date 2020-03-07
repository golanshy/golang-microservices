package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/golanshy/golang-microservices/oauth-api/src/api/domain/oauth"
	"github.com/golanshy/golang-microservices/oauth-api/src/api/services"
	"github.com/golanshy/golang-microservices/src/api/utils/errors"
	"net/http"
)

func CreateAccessToken(c *gin.Context) {
	var request oauth.AccessTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		// request does not contain a valid json body
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	// We have a valid json in the request
	token, err := services.OauthService.CreateAccessToken(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, token)
}

func GetAccessToken(c *gin.Context) {
	token, err := oauth.GetAccessTokenByToken(c.Param("token_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, token)
}
