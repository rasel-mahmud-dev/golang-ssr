package services

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type JwtClaims struct {
	UserId int32 `json:"user_id"`
	jwt.StandardClaims
}

func Gen2(userId string) (string, error) {
	JwtSecret := os.Getenv("JWT_SECRET")
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(10 * time.Hour).Unix()
	claims["iat"] = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)

	tokenString, err := token.SignedString([]byte(JwtSecret))
	return tokenString, err
}

func GenerateToken(userId int32) string {
	JwtSecret := os.Getenv("JWT_SECRET")
	mySigningKey := []byte(JwtSecret)

	expirationTime := time.Now().Add(60 * time.Minute)

	claims := JwtClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "test",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return ""
	}
	return ss
}

func ParseJwtToken(token string) (error, JwtClaims) {
	JwtSecret := os.Getenv("JWT_SECRET")
	tokenExtract, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(JwtSecret), nil
	})

	if err != nil {
		return err, JwtClaims{}
	}

	claims, ok := tokenExtract.Claims.(JwtClaims)

	if ok && tokenExtract.Valid {
		return nil, claims
	}

	return errors.New("jwt expired"), JwtClaims{}
}
