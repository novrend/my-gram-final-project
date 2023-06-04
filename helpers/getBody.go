package helpers

import "github.com/gin-gonic/gin"

var (
	appJSON = "application/json"
)

func GetBody(c *gin.Context, entity interface{}) error {
	var err error
	contentType := c.Request.Header.Get("Content-Type")
	if contentType == appJSON {
		err = c.ShouldBindJSON(&entity)

	} else {
		err = c.ShouldBind(&entity)
	}
	return err
}
