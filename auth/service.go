package auth

import (
	"project/Database/models"
)

func CreateAccountService(request models.User) (models.User, error) {
	account, err := storeAccountRow(request)
	if err != nil {
		return models.User{}, err
	}
	return account, nil
}

func GetUser(email string) (models.User, error) {
	result, err := getUser(email)
	if err != nil {
		return models.User{}, err
	}
	return result, nil
}
