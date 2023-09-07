package main

import (
	"awesomeProject/src/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	routes.Route(route.Group("/api/v1"))

	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
