package main

import (
	"awesomeProject/src/database"
	"awesomeProject/src/initializer"
	"awesomeProject/src/routes"
)

func init() {
	initializer.LoadEnv()
	_, err := database.DbConnect()
	if err != nil {
		panic(err)
	}
}

func main() {
	r := routes.Route()
	r.Run(":8089") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
