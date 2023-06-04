package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `gorm:"column:title;not null" json:"title" form:"title" valid:"required"`
	Caption  string `gorm:"column:caption" json:"caption" form:"caption"`
	PhotoUrl string `gorm:"column:photo_url;not null" json:"photo_url" form:"photo_url" valid:"required"`
	UserId   uint   `gorm:"column:user_id;" json:"user_id" form:"user_id"`
	User     *User  `json:"-"`
}

type UpdatePhotoRequest struct {
	Title    string `json:"title" form:"title"`
	Caption  string `json:"caption" form:"title"`
	PhotoUrl string `json:"photo_url" form:"photo_url"`
}

func (photo *Photo) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(photo)
	if err != nil {
		return err
	}

	return nil
}

func (photo *Photo) BeforeUpdate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(photo)
	if err != nil {
		return err
	}

	return nil
}
