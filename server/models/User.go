package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/vesarian/cms-cloth/helpers"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"column:username;not null;uniqueIndex" json:"username" form:"username" valid:"required"`
	Email    string `gorm:"column:email;not null;uniqueIndex" json:"email" form:"email" valid:"required,email"`
	Password string `gorm:"column:password;not null" json:"password" form:"password" valid:"required"`
	Role     string `gorm:"column:role;not null" json:"role" form:"role" valid:"required"`
}

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}

	if len(user.Password) < 5 {
		return errors.New("password must be more than 5 characters")
	}

	user.Password = helpers.HashPass(user.Password)

	return nil
}
