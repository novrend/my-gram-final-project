package controllers

import (
	"final-project/exceptions"
	"final-project/helpers"
	"final-project/models"
	"final-project/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {

	user := models.User{}

	_ = helpers.GetBody(c, &user)

	err := repositories.CreateUser(&user)

	if err != nil {
		exceptions.NewBadRequestError(c, err.Error())
		return
	}

	response := models.RegisterResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    response,
	})
}

func UserLogin(c *gin.Context) {

	user := models.User{}

	_ = helpers.GetBody(c, &user)

	password := user.Password

	err := repositories.FindUser(&user)

	if err != nil {
		exceptions.NewUnauthorizedError(c, "Invalid email/password")
		return
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))

	if !comparePass {
		exceptions.NewUnauthorizedError(c, "Invalid email/password")
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email, user.Username)

	response := models.LoginResponse{
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})

}
