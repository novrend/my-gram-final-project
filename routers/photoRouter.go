package routers

import (
	"final-project/controllers"
	"final-project/middlewares"
	"final-project/models"
	"github.com/gin-gonic/gin"
)

func PhotoRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetPhotos)
	router.POST("/", controllers.CreatePhoto)
	router.GET("/:photoId", controllers.GetPhotoByID)
	router.Use(middlewares.Authorization("photo", func() interface{} {
		return &models.Photo{}
	}))
	router.PUT("/:photoId", controllers.UpdatePhoto)
	router.DELETE("/:photoId", controllers.DeletePhoto)
}
