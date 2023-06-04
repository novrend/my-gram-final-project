package exceptions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewInternalServerError(c *gin.Context, error string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"error":   "Internal Server Error",
		"message": error,
	})
}
