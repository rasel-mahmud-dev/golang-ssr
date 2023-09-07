package services

import (
	"awesomeProject/src/models"
)

func GetArticles() []models.Article {
	return models.GetArticles()
}

func AddArticle(article models.Article) {
	article.Slug = article.Title
	article.AuthorID = 1
	models.AddArticle(article)
}
