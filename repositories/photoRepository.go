package repositories

import (
	"final-project/database"
	"final-project/models"
)

func GetPhotos(photo *[]models.Photo) error {
	db := database.GetDB()
	err := db.Model(&models.Photo{}).Find(&photo).Error
	if err != nil {
		return err
	}
	return nil
}

func CreatePhoto(photo *models.Photo) error {
	db := database.GetDB()
	err := db.Debug().Create(&photo).Error
	return err
}

func GetPhotoByID(photoId string, photo *models.Photo) error {
	db := database.GetDB()
	result := db.Debug().Where("id = ?", photoId).First(&photo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdatePhoto(photo *models.Photo) error {
	db := database.GetDB()
	err := db.Debug().Save(&photo).Error
	if err != nil {
		return err
	}
	return nil
}

func DeletePhoto(photoId string) error {
	db := database.GetDB()
	err := db.Debug().Delete(&models.Photo{}, photoId).Error
	if err != nil {
		return err
	}
	return nil
}
