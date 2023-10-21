package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rasel-mahmud-dev/golang-ssr/src/controllers"
)

func ArticleRoute(r *gin.RouterGroup) {
	r.GET("/", controllers.GetArticles)
	r.POST("/new", controllers.AddArticle)
}
