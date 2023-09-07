package main

import (
	"awesomeProject/models"
	"database/sql"
	"fmt"
)

func DbConnect() *sql.DB {
	// Connect to the database
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	return db
}

func main() {

	var post = models.InitPost()

	post.AddPost(map[string]interface{}{"title": "Test post s", "price": 22.23})
	post.AddPost(map[string]interface{}{"title": "Test post s", "price": 22.23})

	fmt.Println(post.GetPosts())

	fmt.Println(post.FindAll())

}
