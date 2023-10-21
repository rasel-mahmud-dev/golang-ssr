package services

import (
	"github.com/golang-jwt/jwt"
	"os"
)

type JwtClaims struct {
	UserId int32 `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userId int32) string {
	JwtSecret := os.Getenv("DB_URI")
	mySigningKey := []byte(JwtSecret)

	claims := JwtClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return ""
	}
	return ss
}
