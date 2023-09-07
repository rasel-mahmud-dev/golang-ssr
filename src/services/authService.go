package services

import (
	"awesomeProject/src/models"
)

func Login(email string, password string) map[string]interface{} {
	user := models.FindUser(email)
	return user
}
