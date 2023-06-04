package repositories

import (
	"final-project/database"
	"final-project/models"
)

func GetSocialMedia(socialMedia *[]models.SocialMedia) error {
	db := database.GetDB()
	err := db.Model(&models.SocialMedia{}).Find(&socialMedia).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateSocialMedia(socialMedia *models.SocialMedia) error {
	db := database.GetDB()
	err := db.Debug().Create(&socialMedia).Error
	return err
}

func GetSocialMediaByID(socialMediaId string, socialMedia *models.SocialMedia) error {
	db := database.GetDB()
	err := db.Debug().Where("id = ?", socialMediaId).First(&socialMedia)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func UpdateSocialMedia(socialMedia *models.SocialMedia) error {
	db := database.GetDB()
	err := db.Debug().Save(&socialMedia).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteSocialMedia(socialMediaId string) error {
	db := database.GetDB()
	err := db.Debug().Delete(&models.SocialMedia{}, socialMediaId).Error
	if err != nil {
		return err
	}
	return nil
}
