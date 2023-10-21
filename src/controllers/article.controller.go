package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rasel-mahmud-dev/golang-ssr/src/models"
	"github.com/rasel-mahmud-dev/golang-ssr/src/services"
	"net/http"
)

func GetArticles(c *gin.Context) {
	articles := services.GetArticles()
	c.JSON(200, articles)
}

func AddArticle(c *gin.Context) {
	article := models.Article{}

	err := c.BindJSON(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You can now use the 'article' struct with the request data
	fmt.Printf("Received article: %+v\n", article)

	services.AddArticle(article)
	c.JSON(200, "Message")
}
