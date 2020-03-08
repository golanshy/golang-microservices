package app

import (
	"github.com/golanshy/golang-microservices/go-common/src/api/controllers/polo"
	"github.com/golanshy/golang-microservices/oauth-api/controllers/repositories"
)

func mapUrls() {
	router.POST("/session_token/", repositories.CreateSessionToken)
	router.GET("/marco", polo.Marco)
}
