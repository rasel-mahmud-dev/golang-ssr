package routes

import "github.com/gin-gonic/gin"

func Route() *gin.Engine {

	route := gin.Default()
	route.Group("/api/v1")

	ArticleRoute(route.Group("/articles"))
	AuthRoute(route.Group("/users"))

	return route
}
