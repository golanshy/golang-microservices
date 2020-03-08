package app

import (
	"github.com/golanshy/golang-microservices/demo/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
