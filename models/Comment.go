package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	UserId  uint   `gorm:"column:user_id" json:"user_id" form:"user_id"`
	User    *User  `json:"-"`
	PhotoId uint   `gorm:"column:photo_id" json:"photo_id" form:"photo_id"`
	Photo   *Photo `json:"-"`
	Message string `gorm:"column:message;not null" json:"message" form:"message" valid:"required"`
}

type UpdateCommentRequest struct {
	Message string `json:"message" json:"message"`
}

func (comment *Comment) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(comment)
	if err != nil {
		return err
	}

	return nil
}

func (comment *Comment) BeforeUpdate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(comment)
	if err != nil {
		return err
	}

	return nil
}
