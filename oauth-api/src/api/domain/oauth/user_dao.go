package oauth

import (
	"github.com/golanshy/golang-microservices/src/api/utils/errors"
)

const (
	queryGetUsernameByUsernameAndPassword = "Select id, usernameFROm users WHERE username=? ND password=?;"
)

var (
	users = map[string]*User{
		"golanshy": &User{
			Id:       123,
			Username: "golanshy",
		},
	}
)

func GetUsernameByUsernameAndPassword(username string, password string) (*User, errors.APiError) {
	user := users[username]
	if user == nil {
		return nil, errors.NewNotFoundApiError("no user found with given parameters")
	}
	return user, nil
}
