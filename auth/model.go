package auth

import (
	"project/Database"
	"project/Database/models"
)

func storeAccountRow(row models.User) (models.User, error) {
	db := Database.DB
	err := db.Save(&row).Error
	return row, err
}

func getUser(email string) (models.User, error) {
	db := Database.DB
	var user models.User
	err := db.Where("email = ?", email).First(&user).Error
	return user, err
}
