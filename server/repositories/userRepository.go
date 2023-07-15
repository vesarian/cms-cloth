package repositories

import (
	"github.com/vesarian/cms-cloth/database"
	"github.com/vesarian/cms-cloth/models"
)

func CreateUser(User *models.User)	error {
	db := database.GetDB()
	err := db.Debug().Create(&User).Error
	return err
}

func FindUser(User *models.User) error {
	db := database.GetDB()
	err := db.Debug().Where("email = ? or username= ?", User.Email, User.Username).Take(&User).Error
	return err
}