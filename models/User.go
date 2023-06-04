package models

import (
	"errors"
	"final-project/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"column:username;not null;uniqueIndex" json:"username" form:"username" valid:"required"`
	Email    string `gorm:"column:email;not null;uniqueIndex" json:"email" form:"email" valid:"required,email"`
	Password string `gorm:"column:password;not null" json:"password" form:"password" valid:"required"`
	Age      uint   `gorm:"column:age;not null" json:"age" form:"age" valid:"required,numeric"`
}

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}

	if user.Age < 9 {
		return errors.New("age must be greater than or equal to 9")
	}

	if len(user.Password) < 6 {
		return errors.New("password must have a minimum length of 6 characters")
	}

	user.Password = helpers.HashPass(user.Password)

	return nil
}
