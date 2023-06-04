package repositories

import (
	"final-project/database"
	"final-project/models"
)

func CreateUser(User *models.User) error {
	db := database.GetDB()
	err := db.Debug().Create(&User).Error
	return err
}

func FindUser(User *models.User) error {
	db := database.GetDB()
	err := db.Debug().Where("email = ? OR username = ?", User.Email, User.Username).Take(&User).Error
	return err
}
