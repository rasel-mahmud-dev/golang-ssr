package controllers

import (
	"awesomeProject/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	type Body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var body Body
	err := c.Bind(&body)
	if err != nil {
		return
	}

	user := services.Login(body.Email, body.Password)
	c.JSON(200, user)
}
