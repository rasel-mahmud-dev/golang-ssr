package services

import "awesomeProject/models"

func Login(email string, password string) map[string]interface{} {
	userRepo := models.InitUserRepository()
	user := userRepo.FindUser(email)
	return user
}
