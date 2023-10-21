package services

import (
	"errors"
	"fmt"
	"github.com/rasel-mahmud-dev/golang-ssr/src/models"
	"github.com/rasel-mahmud-dev/golang-ssr/src/repository"
	"github.com/rasel-mahmud-dev/golang-ssr/src/services/hash"
	"strconv"
)

func Login(email string, password string) (interface{}, error) {
	user, _ := repository.FindUser(email)

	if !hash.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("password wrong .")
	}

	token := GenerateToken(int32(user.ID))

	return map[string]interface{}{
		"user": map[string]interface{}{
			"id":    user.ID,
			"email": user.Email,
		},
		"token": token,
	}, nil
}

func CreateUser(user models.User) (interface{}, error) {
	result, err := repository.CreateUser(user)
	return result, err

}

func VerifyAuthUser(token string) (interface{}, error) {

	//t := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE2OTc5MTA4MTUsImlhdCI6MTY5NzkwNzIxNSwiaXNzIjoidGVzdCJ9.r5ahSiyutS65kVYaeWgpCH2y395xrSyryqQGUY0bYBQ"

	aa, err := Gen2(strconv.Itoa(234))

	err, data := ParseJwtToken(aa)
	fmt.Println(data.UserId)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"user": "userInfo",
	}, nil
}
