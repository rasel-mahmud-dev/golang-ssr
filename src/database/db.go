package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var Db *sql.DB

func DbConnect() (*sql.DB, error) {
	dns := os.Getenv("DB_URI")
	var err error = nil
	Db, err = sql.Open("postgres", dns)
	if err != nil {
		fmt.Println("Database connection fail!")
		err = errors.New("Database connection fail!")
		return Db, err
	}

	fmt.Println("Connected to the database!")

	return Db, err

}
