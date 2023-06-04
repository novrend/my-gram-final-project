package exceptions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewNotFoundError(c *gin.Context, error string) {
	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"error":   "Data Not Found",
		"message": error,
	})
}
