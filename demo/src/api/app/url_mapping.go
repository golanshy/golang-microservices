package app

import (
	"github.com/golanshy/golang-microservices/demo/src/api/controllers/repositories"
	"github.com/golanshy/golang-microservices/go-common/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
