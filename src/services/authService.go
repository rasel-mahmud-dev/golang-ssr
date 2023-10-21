package services

import (
	"errors"
	"github.com/rasel-mahmud-dev/golang-ssr/src/models"
)

func Login(email string, password string) (interface{}, error) {
	user, _ := models.FindUser(email)

	if user.Password != password {
		return nil, errors.New("password wrong")
	}

	return map[string]interface{}{
		"user": map[string]interface{}{
			"id":    user.ID,
			"email": user.Email,
		},
		"token": "generated jwt token..",
	}, nil
}
