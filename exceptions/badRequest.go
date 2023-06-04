package exceptions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewBadRequestError(c *gin.Context, error string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   "Bad Request",
		"message": error,
	})
}
