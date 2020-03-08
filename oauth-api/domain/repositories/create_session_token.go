package repositories

import (
	"github.com/golanshy/golang-microservices/go-common/utils/errors"
	"strings"
)

type CreateSessionTokenRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrandType  string `json:"grant_type"`
	Audience  string `json:"audience"`
}

func (r *CreateSessionTokenRequest) Validate() errors.APiError {
	r.ClientId = strings.TrimSpace(r.ClientId)
	if r.ClientId == "" {
		return errors.NewBadRequestError("Invalid ClientId")
	}
	r.ClientSecret = strings.TrimSpace(r.ClientSecret)
	if r.ClientSecret == "" {
		return errors.NewBadRequestError("Invalid Client Secret")
	}
	return nil
}

type CreateSessionTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
