package services

import (
	"github.com/golanshy/golang-microservices/oauth-api/src/api/domain/oauth"
	"github.com/golanshy/golang-microservices/src/api/utils/errors"
	"time"
)

type oauthService struct {
}

type oauthServiceInterface interface {
	CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.APiError)
	GetAccessToken(accessToken string) (*oauth.AccessToken, errors.APiError)
}

var OauthService oauthServiceInterface

func init() {
	OauthService = &oauthService{}
}

func (s *oauthService) CreateAccessToken(request oauth.AccessTokenRequest) (*oauth.AccessToken, errors.APiError) {
	if err:= request.Validate(); err != nil {
		return nil, err
	}
	// We have a valid username & password
	user, err := oauth.GetUsernameByUsernameAndPassword(request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	token := oauth.AccessToken{
		UserId: user.Id,
		Expires: time.Now().UTC().Add(24 * time.Hour).Unix(),
	}

	if err:= token.Save(); err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *oauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.APiError) {
	token, err := oauth.GetAccessTokenByToken(accessToken)
	if err != nil {
		return nil, err
	}
	if token.IsExpired() {
		return nil, errors.NewNotFoundApiError("no access token found for given parameters")
	}
	return token, nil
}
