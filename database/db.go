package database

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"os"
)

func DbConnect() *sql.DB {
	// Connect to the database
	db, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	file, err := os.ReadFile("seed_sql.sql")
	if err != nil {
		return nil
	}

	_, err = db.Exec(string(file))
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
