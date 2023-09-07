package routes

import (
	"awesomeProject/src/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.RouterGroup) {
	r.POST("/login", controllers.Login)
}
