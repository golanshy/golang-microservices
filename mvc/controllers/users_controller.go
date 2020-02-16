package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golanshy/golang-microservices/mvc/services"
	"github.com/golanshy/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		// Just return the bad request to the client
		apiErr:= &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode:  http.StatusBadRequest,
			Code:    "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr!= nil {
		// Handle error and return to client
		utils.RespondError(c, apiErr)
		return
	}

	// Return user to client
	utils.Respond(c, http.StatusOK, user)
}

