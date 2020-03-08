package services

import (
	"fmt"
	"github.com/golanshy/golang-microservices/go-common/log/logrus"
	"github.com/golanshy/golang-microservices/go-common/log/zap"
	"github.com/golanshy/golang-microservices/go-common/utils/errors"
	"github.com/golanshy/golang-microservices/oauth-api/domain/repositories"
	"github.com/golanshy/golang-microservices/oauth-api/providers"
)

type repoService struct {
}

type repoServiceInterface interface {
	CreateSessionToken(request repositories.CreateSessionTokenRequest) (*repositories.CreateSessionTokenResponse, errors.APiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateSessionToken(input repositories.CreateSessionTokenRequest) (*repositories.CreateSessionTokenResponse, errors.APiError) {

	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := repositories.CreateSessionTokenRequest{
		ClientId:        input.ClientId,
		ClientSecret: input.ClientSecret,
		GrandType: "client_credentials",
		Audience: "https://applylogic.auth0.com/api/v2/",
	}
	logrus.Info("about to send request to external api",
		fmt.Sprintf("clientId:%s", input.ClientId),
		"status:pending")
	zap.Info("about to send request to external api",
		zap.Field("clientId", input.ClientId),
		zap.Field("status", "pending"),
		zap.Field("authenticated", input.ClientId != ""))

	response, err := providers.CreateSessionToken(request)
	if err != nil {
		logrus.Error("response obtained from external api",
			err.Error,
			fmt.Sprintf("clientId:%s", input.ClientId),
			"status:error")
		zap.Error("response obtained from external api",
			err.Error,
			zap.Field("error", err.Error),
			zap.Field("errorDescription", err.ErrorDescription),
			zap.Field("statusCode", err.StatusCode),
			zap.Field("clientId", input.ClientId),
			zap.Field("status", "error"),
			zap.Field("authenticated", input.ClientId != ""))

		return nil, errors.NewApiError(err.StatusCode, err.ErrorDescription)
	}
	logrus.Info("response obtained from external api",
		fmt.Sprintf("clientId:%s", input.ClientId),
		"status:success")
	zap.Info("response obtained from external api",
		zap.Field("clientId", input.ClientId),
		zap.Field("status", "success"),
		zap.Field("authenticated", input.ClientId != ""))

	result := &repositories.CreateSessionTokenResponse{
		AccessToken:    response.AccessToken,
		TokenType:  response.TokenType,
	}

	return result, nil
}
