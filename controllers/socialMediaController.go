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

func GetSocialMedia(c *gin.Context) {
	var socialMedia []models.SocialMedia
	err := repositories.GetSocialMedia(&socialMedia)
	if err != nil {
		exceptions.NewInternalServerError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    socialMedia,
	})
}

func CreateSocialMedia(c *gin.Context) {
	socialMedia := models.SocialMedia{}

	_ = helpers.GetBody(c, &socialMedia)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	socialMedia.UserId = userID

	err := repositories.CreateSocialMedia(&socialMedia)
	if err != nil {
		exceptions.NewBadRequestError(c, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    socialMedia,
	})
}

func GetSocialMediaByID(c *gin.Context) {
	socialMedia := models.SocialMedia{}

	_ = helpers.GetBody(c, &socialMedia)

	socialMediaId := c.Param("socialMediaId")

	err := repositories.GetSocialMediaByID(socialMediaId, &socialMedia)
	if err != nil {
		exceptions.NewNotFoundError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    socialMedia,
	})
}

func UpdateSocialMedia(c *gin.Context) {
	socialMediaData := models.UpdateSocialMediaRequest{}

	_ = helpers.GetBody(c, &socialMediaData)

	socialMedia := c.MustGet("socialMedia").(*models.SocialMedia)
	socialMedia.SocialMediaUrl = socialMediaData.SocialMediaUrl
	socialMedia.Name = socialMediaData.Name

	err := repositories.UpdateSocialMedia(socialMedia)
	if err != nil {
		exceptions.NewBadRequestError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    &socialMedia,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	socialMediaId := c.Param("socialMediaId")

	err := repositories.DeleteSocialMedia(socialMediaId)
	if err != nil {
		exceptions.NewBadRequestError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    "",
	})
}
