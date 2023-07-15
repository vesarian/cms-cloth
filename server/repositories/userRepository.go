package repositories

import (
	"github.com/vesarian/cms-cloth/database"
	"github.com/vesarian/cms-cloth/models"
)

func GetUser(User *[]models.User) error {
	db := database.GetDB()
	err := db.Model(&models.User{}).Find(&User).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(User *models.User) error {
	db := database.GetDB()
	err := db.Debug().Create(&User).Error
	return err
}

func FindUser(User *models.User) error {
	db := database.GetDB()
	err := db.Debug().Where("email = ? or username= ?", User.Email, User.Username).Take(&User).Error
	return err
}
