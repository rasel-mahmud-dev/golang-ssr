package main

import (
	"fmt"
	"github.com/rasel-mahmud-dev/golang-ssr/src/database"
	"github.com/rasel-mahmud-dev/golang-ssr/src/initializer"
	"github.com/rasel-mahmud-dev/golang-ssr/src/services"
)

func init() {
	initializer.LoadEnv()
	_, err := database.DbConnect()
	if err != nil {
		panic(err)
	}
}

func main() {
	//r := routes.Route()
	//err := r.Run(":8089")
	//if err != nil {
	//	fmt.Println("Server error..")
	//	return
	//} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	token := services.GenerateToken(234)
	fmt.Println(token)
	err, j := services.ParseJwtToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(j)

}
