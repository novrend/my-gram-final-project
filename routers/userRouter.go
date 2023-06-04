package routers

import (
	"final-project/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	router.POST("/register", controllers.UserRegister)
	router.POST("/login", controllers.UserLogin)
}
