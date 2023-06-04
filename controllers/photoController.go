package controllers

import (
	"final-project/exceptions"
	"final-project/helpers"
	"final-project/models"
	"final-project/repositories"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	err := repositories.GetPhotos(&photos)
	if err != nil {
		exceptions.NewInternalServerError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    photos,
	})
}

func CreatePhoto(c *gin.Context) {
	photo := models.Photo{}

	_ = helpers.GetBody(c, &photo)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	photo.UserId = userID

	err := repositories.CreatePhoto(&photo)
	if err != nil {
		exceptions.NewBadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    photo,
	})
}

func GetPhotoByID(c *gin.Context) {
	photo := models.Photo{}

	_ = helpers.GetBody(c, &photo)

	photoId := c.Param("photoId")

	err := repositories.GetPhotoByID(photoId, &photo)
	if err != nil {
		exceptions.NewNotFoundError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    photo,
	})
}

func UpdatePhoto(c *gin.Context) {
	photoData := models.UpdatePhotoRequest{}

	_ = helpers.GetBody(c, &photoData)

	photo := c.MustGet("photo").(*models.Photo)
	photo.Title = photoData.Title
	photo.Caption = photoData.Caption
	photo.PhotoUrl = photoData.PhotoUrl

	err := repositories.UpdatePhoto(photo)
	if err != nil {
		exceptions.NewBadRequestError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    &photo,
	})
}

func DeletePhoto(c *gin.Context) {
	photoId := c.Param("photoId")

	err := repositories.DeletePhoto(photoId)
	if err != nil {
		exceptions.NewBadRequestError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    "",
	})
}
