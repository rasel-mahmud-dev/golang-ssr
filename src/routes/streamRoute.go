package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rasel-mahmud-dev/golang-ssr/src/controllers"
)

func StreamRoute(r *gin.RouterGroup) {
	r.GET("/:fileName", controllers.Stream2)
}
