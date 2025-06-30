package Database

import (
	"fmt"
	"os"

	"project/Database/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	dsn := os.Getenv("CONN_STR")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to the database: %v", err)
	}

	err = DB.AutoMigrate(&models.Brand{}, &models.Product{}, &models.Tag{},&models.User{})

	if err != nil {
		fmt.Println("Faild To Migrate. ", err)
	}
}
