package routes

import "github.com/gin-gonic/gin"

func Route(r *gin.RouterGroup) {
	ArticleRoute(r.Group("/articles"))
	AuthRoute(r.Group("/users"))
}
