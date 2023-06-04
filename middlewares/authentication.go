package middlewares

import (
	"final-project/exceptions"
	"final-project/helpers"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			exceptions.NewUnauthenticatedError(c)
			c.Abort()
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}
