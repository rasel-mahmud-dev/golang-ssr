package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rasel-mahmud-dev/golang-ssr/src/controllers"
)

func AuthRoute(r *gin.RouterGroup) {
	r.GET("users", controllers.GetAllUsers)
	r.POST("login", controllers.Login)
	r.POST("registration", controllers.Registration)
	r.GET("verify", controllers.AuthVerify)
}
