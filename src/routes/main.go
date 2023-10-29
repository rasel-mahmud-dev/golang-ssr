package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rasel-mahmud-dev/golang-ssr/src/models"
	"net/http"
)

func Route() *gin.Engine {

	route := gin.Default()

	api := route.Group("/api")

	ArticleRoute(api.Group("/articles"))
	AuthRoute(api.Group("/auth"))
	StreamRoute(api.Group("/stream"))

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

	route.GET("/add-article", func(context *gin.Context) {
		context.HTML(http.StatusOK, "add-article.gohtml", nil)
	})

	return route
}
