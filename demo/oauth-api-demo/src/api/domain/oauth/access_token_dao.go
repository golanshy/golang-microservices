package oauth

import (
	"fmt"
	"github.com/golanshy/golang-microservices/go-common/utils/errors"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

func (at *AccessToken) Save() errors.APiError {
	at.AccessToken = fmt.Sprintf("USR_%d", at.UserId)
	tokens[at.AccessToken] = at
	return nil
}

func GetAccessTokenByToken(accessToken string) (*AccessToken, errors.APiError) {
	token := tokens[accessToken]
	if token == nil {
		return nil, errors.NewNotFoundApiError("no access token found with given parameters")
	}
	return token, nil
}
