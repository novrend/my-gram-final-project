package exceptions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewUnauthorizedError(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"error":   "Unauthorized",
		"message": message,
	})
}
