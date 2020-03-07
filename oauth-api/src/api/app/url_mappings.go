package app

import (
	"github.com/golanshy/golang-microservices/oauth-api/src/api/controlers/oauth"
	"github.com/golanshy/golang-microservices/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)

}