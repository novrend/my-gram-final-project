package routers

import (
	"final-project/middlewares"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	usersRouter := router.Group("/users")
	UserRoutes(usersRouter)

	router.Use(middlewares.Authentication())

	photoRouter := router.Group("/photos")
	PhotoRoutes(photoRouter)

	commentRouter := router.Group("/comments")
	CommentRoutes(commentRouter)

	socialMediaRouter := router.Group("/social-media")
	SocialMediaRoutes(socialMediaRouter)

	return router
}
