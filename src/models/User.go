package models

import (
	"time"
)

type User struct {
	ID                  int64     `json:"id"`
	FirstName           string    `json:"first_name"`
	LastName            string    `json:"last_name"`
	Email               string    `json:"email"`
	Password            string    `json:"password"`
	ProfilePictureURL   string    `json:"profile_picture_url"`
	CreatedAt           time.Time `json:"created_at"`
	IsVerified          bool      `json:"is_verified"`
	VerificationCode    int64     `json:"verification_code"`
	VerificationExpired time.Time `json:"verification_expired"`
}

func AddUser(article Article) error {
	//_, err := database.Db.Exec(`INSERT INTO articles (title, content, author_id, slug, cover_image_url, status, word_count, read_time)
	//                   VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
	//	article.Title, article.Content, article.AuthorID, article.Slug, article.CoverImageURL, article.Status, article.WordCount, article.ReadTime)
	//
	//if err != nil {
	//	fmt.Println(
	//		err)
	//	return err
	//}
	return nil
}
