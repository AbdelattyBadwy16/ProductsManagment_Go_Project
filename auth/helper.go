package auth

import "project/Database/models"

func setAccount(request RegisterRequest) models.User {
	return models.User{
		Name: request.Name,
		Password : request.Password,
		Email : request.Email,
	}
}
