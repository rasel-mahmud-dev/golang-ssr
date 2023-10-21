package services

import (
	"errors"
	"fmt"
	"github.com/rasel-mahmud-dev/golang-ssr/src/models"
	"strconv"
)

func Login(email string, password string) (interface{}, error) {
	user, _ := models.FindUser(email)

	if user.Password != password {
		return nil, errors.New("password wrong")
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
