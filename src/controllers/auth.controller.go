package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rasel-mahmud-dev/golang-ssr/src/models"
	"github.com/rasel-mahmud-dev/golang-ssr/src/repository"
	"github.com/rasel-mahmud-dev/golang-ssr/src/services"
	"net/http"
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
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, user)
}

func GetAllUsers(c *gin.Context) {
	users, err := repository.FindUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func Registration(c *gin.Context) {
	var user models.User

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	result, err := services.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, result)
}

func AuthVerify(c *gin.Context) {
	authorization := c.GetHeader("authorization")
	parts := strings.Split(authorization, " ")
	token := parts[1]

	user, err := services.VerifyAuthUser(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, user)
}
