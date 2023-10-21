package repository

import (
	"errors"
	"fmt"
	"github.com/rasel-mahmud-dev/golang-ssr/src/database"
	"github.com/rasel-mahmud-dev/golang-ssr/src/models"
	"github.com/rasel-mahmud-dev/golang-ssr/src/services/hash"
	"strings"
)

func FindUsers() ([]models.User, error) {
	var users []models.User

	rows, err := database.Db.Query(`select id, email, first_name, last_name, password from users`)
	if err != nil {
		return users, nil
	}

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func FindUser(email string) (models.User, error) {
	var user models.User
	err := database.Db.QueryRow(`select id, email, password from users  u where u.email = $1`, email).
		Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		fmt.Println(err.Error())
		return user, nil
	}
	return user, nil
}

func CreateUser(user models.User) (int64, error) {
	hashPass, err := hash.HashPassword(user.Password)

	if err != nil {
		return 0, errors.New("Password hash fail")
	}

	sql := `
		insert into users (
		   first_name, 
		   last_name,
		   email,
		   password
		) values ($1, $2, $3, $4)
	`
	result, err := database.Db.Exec(
		sql,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		hashPass,
	)

	if err != nil {
		msg := err.Error()
		hasAcc := strings.Contains(msg, "duplicate key value violates unique constraint \"users_email_key\"")
		if hasAcc {
			return 0, errors.New("This email is already registered")
		}
		return 0, err
	}

	id, _ := result.LastInsertId()

	return id, nil

}
