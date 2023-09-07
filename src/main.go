package main

import (
	"awesomeProject/src/database"
	"awesomeProject/src/routes"
)

func main() {
	database.DbConnect()

	r := routes.Route()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
