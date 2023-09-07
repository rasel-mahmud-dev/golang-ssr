package models

import (
	"awesomeProject/src/database"
	"fmt"
)

type User struct {
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

func AddUser(article Article) error {
	_, err := database.Db.Exec(`INSERT INTO articles (title, content, author_id, slug, cover_image_url, status, word_count, read_time)
                       VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		article.Title, article.Content, article.AuthorID, article.Slug, article.CoverImageURL, article.Status, article.WordCount, article.ReadTime)

	if err != nil {
		fmt.Println(
			err)
		return err
	}
	return nil
}

func FindUser(email string) map[string]interface{} {
	var user map[string]interface{}
	err := database.Db.QueryRow(`select user_id, username, password_hash, email, profile_picture_url from users where email = ?`, email).Scan(&user)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return user
}
