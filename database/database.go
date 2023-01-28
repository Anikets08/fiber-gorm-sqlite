package database

import (
	"log"

	"github.com/anikets08/go-rest/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&models.Restaurant{}, &models.Category{}, &models.MenuItem{})
}
