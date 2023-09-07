package services

import (
	"awesomeProject/src/models"
)

func GetArticles() []models.Article {
	a := models.InitArticleRepository()
	return a.GetArticles()
}

func AddArticle(article models.Article) {
	a := models.InitArticleRepository()
	err := a.AddArticle(article)
	if err != nil {
		return
	}
}
