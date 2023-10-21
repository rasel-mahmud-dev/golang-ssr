package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rasel-mahmud-dev/golang-ssr/src/services"
	"strings"
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

	user, err := services.Login(body.Email, body.Password)

	if err != nil {
		c.JSON(500, map[string]string{"message": err.Error()})
		return
	}

	c.JSON(200, user)

}

func AuthVerify(c *gin.Context) {
	authorization := c.GetHeader("authorization")
	parts := strings.Split(authorization, " ")
	token := parts[1]

	user, err := services.VerifyAuthUser(token)
	if err != nil {
		c.JSON(500, map[string]string{"message": err.Error()})
		return
	}

	c.JSON(200, user)
}
