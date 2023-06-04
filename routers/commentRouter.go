package routers

import (
	"final-project/controllers"
	"final-project/middlewares"
	"final-project/models"
	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetComments)
	router.POST("/:photoId", controllers.CreateComment)
	router.GET("/:photoId", controllers.GetCommentByPhotoID)
	router.Use(middlewares.Authorization("comment", func() interface{} {
		return &models.Comment{}
	}))
	router.PUT("/:commentId", controllers.UpdateComment)
	router.DELETE("/:commentId", controllers.DeleteComment)
}
