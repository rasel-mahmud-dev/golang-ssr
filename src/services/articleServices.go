package services

import (
	"awesomeProject/src/models"
)

func GetArticles() []models.Article {
	return models.GetArticles()
}

func AddArticle(article models.Article) {
	models.AddArticle(article)
}
