package models

import (
	"time"
)

type User struct {
	ID                  int64     `db:"id"`
	FirstName           string    `db:"first_name"`
	LastName            string    `db:"last_name"`
	Email               string    `db:"email"`
	Password            string    `db:"password"`
	ProfilePictureURL   string    `db:"profile_picture_url"`
	CreatedAt           time.Time `db:"created_at"`
	IsVerified          bool      `db:"is_verified"`
	VerificationCode    int64     `db:"verification_code"`
	VerificationExpired time.Time `db:"verification_expired"`
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
