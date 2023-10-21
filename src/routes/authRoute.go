package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rasel-mahmud-dev/golang-ssr/src/controllers"
)

func AuthRoute(r *gin.RouterGroup) {
	r.POST("login", controllers.Login)
}
