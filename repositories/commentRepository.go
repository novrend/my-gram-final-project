package repositories

import (
	"final-project/database"
	"final-project/models"
)

func GetComments(comments *[]models.Comment) error {
	db := database.GetDB()
	err := db.Debug().Model(&models.Comment{}).Find(&comments).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateComment(comment *models.Comment) error {
	db := database.GetDB()
	err := db.Debug().Create(&comment).Error
	return err
}

func GetCommentsByPhotoID(photoId string, comments *[]models.Comment) error {
	db := database.GetDB()
	err := db.Debug().Where("photo_id = ?", photoId).Find(&comments)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func UpdateComment(comment *models.Comment) error {
	db := database.GetDB()
	err := db.Debug().Save(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(commentId string) error {
	db := database.GetDB()
	err := db.Debug().Delete(&models.Comment{}, commentId).Error
	if err != nil {
		return err
	}
	return nil
}
