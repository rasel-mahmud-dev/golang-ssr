package models

import (
	"database/sql"
	"fmt"
)

type Article struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	AuthorID      int    `json:"author_id"`
	Slug          string `json:"slug"`
	CoverImageURL string `json:"cover_image_url"`
	Status        string `json:"status"`
	WordCount     int    `json:"word_count"`
	ReadTime      int    `json:"read_time"`
}

type ArticleRepository struct {
	Common
}

func (a *ArticleRepository) AddArticle(article Article) error {
	_, err := a.Db.Exec(`INSERT INTO articles (title, content, author_id, slug, cover_image_url, status, word_count, read_time) 
                        VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		article.Title, article.Content, article.AuthorID, article.Slug, article.CoverImageURL, article.Status, article.WordCount, article.ReadTime)

	if err != nil {
		fmt.Println(
			err)
		return err
	}
	return nil
}

func InitArticleRepository() *ArticleRepository {
	return &ArticleRepository{
		Common: *DatabaseConnectionInit("articles"),
	}
}

func (a *ArticleRepository) GetArticles() []Article {
	rows, err := a.Db.Query(`select article_id, slug from articles`)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
		}
	}(rows)

	var articles []Article
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.Id, &article.Slug)
		if err != nil {
			return nil
		}
		articles = append(articles, article)
	}
	return articles
}
