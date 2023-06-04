package middlewares

import (
	"final-project/database"
	"final-project/exceptions"
	"final-project/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Authorization(entityType string, getModel func() interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		id, _ := strconv.Atoi(c.Param(entityType + "Id"))

		entity := getModel()
		result := db.Where("id = ?", id).First(entity)
		if result.Error != nil {
			exceptions.NewNotFoundError(c, result.Error.Error())
			c.Abort()
			return
		} else if result.RowsAffected == 0 {
			exceptions.NewUnauthorizedError(c, "Unauthorized")
			c.Abort()
			return
		}

		switch entity := entity.(type) {
		case *models.Photo:
			if entity.UserId != userID {
				exceptions.NewUnauthorizedError(c, "Unauthorized")
				c.Abort()
				return
			}
		case *models.Comment:
			if entity.UserId != userID {
				exceptions.NewUnauthorizedError(c, "Unauthorized")
				c.Abort()
				return
			}
		case *models.SocialMedia:
			if entity.UserId != userID {
				exceptions.NewUnauthorizedError(c, "Unauthorized")
				c.Abort()
				return
			}
		}

		c.Set(entityType, entity)

		c.Next()
	}
}
