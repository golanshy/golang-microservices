package controllers

import (
	"encoding/json"
	"github.com/golanshy/golang-microservices/mvc/services"
	"github.com/golanshy/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		// Just return the bad request to the client
		apiErr:= &utils.ApplicationError{
			Message: "user_id must be a number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.Status)
		resp.Write([]byte(jsonValue))
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr!= nil {
		// Handle error and return to client
		resp.WriteHeader(apiErr.Status)
		jsonValue, _ := json.Marshal(apiErr)
		resp.Write([]byte(jsonValue))
		return
	}

	// Return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}