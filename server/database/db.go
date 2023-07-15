package database

import (
	"fmt"
	"log"

	// "github.com/vesarian/cms-cloth/models"
	"os"

	"github.com/joho/godotenv"
	"github.com/vesarian/cms-cloth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}
	// end connection
	db.AutoMigrate(&models.User{}, models.Category{}, models.Cloth{})
}

func GetDB() *gorm.DB {
	return db
}
