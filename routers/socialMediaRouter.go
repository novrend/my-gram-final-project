package routers

import (
	"final-project/controllers"
	"final-project/middlewares"
	"final-project/models"
	"github.com/gin-gonic/gin"
)

func SocialMediaRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetSocialMedia)
	router.POST("/", controllers.CreateSocialMedia)
	router.GET("/:socialMediaId", controllers.GetSocialMediaByID)
	router.Use(middlewares.Authorization("socialMedia", func() interface{} {
		return &models.SocialMedia{}
	}))
	router.PUT("/:socialMediaId", controllers.UpdateSocialMedia)
	router.DELETE("/:socialMediaId", controllers.DeleteSocialMedia)
}
