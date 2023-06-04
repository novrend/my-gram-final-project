package controllers

import (
	"final-project/exceptions"
	"final-project/helpers"
	"final-project/models"
	"final-project/repositories"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetComments(c *gin.Context) {
	var comments []models.Comment
	err := repositories.GetComments(&comments)
	if err != nil {
		exceptions.NewInternalServerError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    comments,
	})
}

func CreateComment(c *gin.Context) {
	comment := models.Comment{}

	_ = helpers.GetBody(c, &comment)

	photoId := c.Param("photoId")
	photo := models.Photo{}
	err := repositories.GetPhotoByID(photoId, &photo)
	if err != nil {
		exceptions.NewNotFoundError(c, err.Error())
		return
	}

	id, _ := strconv.ParseUint(photoId, 10, 32)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	comment.UserId = userID
	comment.PhotoId = uint(id)

	err = repositories.CreateComment(&comment)
	if err != nil {
		exceptions.NewBadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    comment,
	})
}

func GetCommentByPhotoID(c *gin.Context) {
	var comments []models.Comment

	_ = helpers.GetBody(c, &comments)

	photoId := c.Param("photoId")
	photo := models.Photo{}
	err := repositories.GetPhotoByID(photoId, &photo)
	if err != nil {
		exceptions.NewNotFoundError(c, err.Error())
		return
	}

	err = repositories.GetCommentsByPhotoID(photoId, &comments)
	if err != nil {
		exceptions.NewNotFoundError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    comments,
	})
}

func UpdateComment(c *gin.Context) {
	commentData := models.UpdateCommentRequest{}

	_ = helpers.GetBody(c, &commentData)

	comment := c.MustGet("comment").(*models.Comment)
	comment.Message = commentData.Message

	err := repositories.UpdateComment(comment)
	if err != nil {
		exceptions.NewBadRequestError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    &comment,
	})
}

func DeleteComment(c *gin.Context) {
	commentId := c.Param("commentId")

	err := repositories.DeleteComment(commentId)
	if err != nil {
		exceptions.NewBadRequestError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    "",
	})
}
