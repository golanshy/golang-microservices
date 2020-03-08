package providers

import (
	"encoding/json"
	"fmt"
	"github.com/golanshy/golang-microservices/go-common/src/api/clients/restClient"
	"github.com/golanshy/golang-microservices/oauth-api/domain/repositories"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	contentType               = "content-type"
	contentTypeApplication    = "application/json"
	headerAuthorization = "authorization"
	headerAuthorizationFormat = "Bearer %s"
	urlCreateSessionToken     = "https://applylogic.auth0.com/oauth/token"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateSessionToken(request repositories.CreateSessionTokenRequest) (*repositories.CreateSessionTokenResponse, *repositories.Oauth0ErrorResponse) {
	headers := http.Header{}
	headers.Set(contentType, contentTypeApplication)

	response, err := restClient.Post(urlCreateSessionToken, request, headers)
	fmt.Println(response)
	if err != nil {
		log.Printf("error when trying to create session token: %s", err.Error())
		return nil, &repositories.Oauth0ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			ErrorDescription:      err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &repositories.Oauth0ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "invalid_response_body",
			ErrorDescription:      "invalid response body",
		}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse repositories.Oauth0ErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &repositories.Oauth0ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Error:      "invalid_json_response_body",
				ErrorDescription:      "invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result repositories.CreateSessionTokenResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Printf("error when trying to unmarshall create repo successful response %s", err.Error())
		return nil, &repositories.Oauth0ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "invalid_json_response_body",
			ErrorDescription:      "error when trying to unmarshal create session token response",
		}
	}
	return &result, nil
}
