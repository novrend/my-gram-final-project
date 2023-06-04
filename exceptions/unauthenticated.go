package exceptions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewUnauthenticatedError(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"error":   "Unauthenticated",
		"message": "Unauthenticated",
	})
}
