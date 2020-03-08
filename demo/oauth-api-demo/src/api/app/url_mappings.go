package app

import (
	"github.com/golanshy/golang-microservices/demo/oauth-api-demo/src/api/controlers/oauth"
	"github.com/golanshy/golang-microservices/go-common/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)

}