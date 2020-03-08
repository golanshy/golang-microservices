package oauth

import (
	"github.com/golanshy/golang-microservices/go-common/utils/errors"
	"strings"
)

type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *AccessTokenRequest) Validate() errors.APiError {
	r.Username = strings.TrimSpace(r.Username)
	if r.Username == "" {
		return errors.NewBadRequestError("invalid username")
	}
	if r.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}