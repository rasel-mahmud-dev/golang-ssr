package models

import (
	"fmt"
	"github.com/rasel-mahmud-dev/golang-ssr/src/database"
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

func FindUser(email string) (User, error) {
	var user User
	err := database.Db.QueryRow(`select id, email, password from users  u where u.email = $1`, email).
		Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		fmt.Println(err.Error())
		return user, nil
	}
	return user, nil
}
