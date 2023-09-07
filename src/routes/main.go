package routes

import (
	"awesomeProject/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route() *gin.Engine {

	route := gin.Default()

	route.Static("/static", "src/public")

	// Configure the HTML template rendering
	route.LoadHTMLGlob("src/views/*.gohtml")

	route.GET("/", func(context *gin.Context) {
		articles := models.GetArticles()
		context.HTML(http.StatusOK, "index.gohtml", gin.H{"Articles": articles})
	})

	route.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.gohtml", nil)
	})

	route.Group("/api/v1")

	ArticleRoute(route.Group("/articles"))
	AuthRoute(route.Group("/users"))

	return route
}
