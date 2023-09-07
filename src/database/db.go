package database

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"os"
)

var Db *sql.DB

func DbConnect() {
	// Connect to the database
	cdb, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		panic(err)
	}
	Db = cdb
	//defer db.Close()

	file, err := os.ReadFile("seed_sql.sql")
	if err != nil {
		fmt.Println(err)
	}

	_, err = Db.Exec(string(file))

}
