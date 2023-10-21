package services

import (
	"github.com/rasel-mahmud-dev/golang-ssr/src/models"
)

func GetArticles() []models.Article {
	return models.GetArticles()
}

func AddArticle(article models.Article) {
	article.Slug = article.Title
	article.AuthorID = 1
	models.AddArticle(article)
}
