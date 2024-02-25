package database

import (
	"os"
	"pp-points/database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Database() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_STRING")))

	if err != nil {
		panic("Database was not able to connect")
	}

	db.AutoMigrate(&models.Groups{})
	db.AutoMigrate(&models.Teachers{})
	db.AutoMigrate(&models.Students{})
	db.AutoMigrate(models.Users{})

	Db = db
}