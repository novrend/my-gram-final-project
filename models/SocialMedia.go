package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"column:name;not null" json:"name" form:"name" valid:"required"`
	SocialMediaUrl string `gorm:"column:social_media_url;not null" json:"social_media_url" form:"social_media_url" valid:"required"`
	UserId         uint   `gorm:"column:user_id" json:"user_id" form:"user_id"`
	User           *User  `json:"-"`
}

type UpdateSocialMediaRequest struct {
	Name           string `json:"name" form:"name"`
	SocialMediaUrl string `json:"social_media_url" form:"social_media_url"`
}

func (socialMedia *SocialMedia) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(socialMedia)
	if err != nil {
		return err
	}

	return nil
}

func (socialMedia *SocialMedia) BeforeUpdate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(socialMedia)
	if err != nil {
		return err
	}

	return nil
}
