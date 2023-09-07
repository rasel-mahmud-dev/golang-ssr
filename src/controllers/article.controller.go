package controllers

import (
	"awesomeProject/src/models"
	"awesomeProject/src/services"
	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	articles := services.GetArticles()
	c.JSON(200, articles)
}

func AddArticle(c *gin.Context) {
	article := models.Article{
		Id:            "12",
		Title:         "",
		Content:       "",
		AuthorID:      0,
		Slug:          "",
		CoverImageURL: "",
		Status:        "",
		WordCount:     0,
		ReadTime:      0,
	}
	services.AddArticle(article)
	c.JSON(200, "Message")
}
