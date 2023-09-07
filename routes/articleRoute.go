package routes

import (
	"awesomeProject/controllers"
	"github.com/gin-gonic/gin"
)

func ArticleRoute(r *gin.RouterGroup) {
	r.GET("/", controllers.GetArticles)
	r.POST("/new", controllers.AddArticle)
}
