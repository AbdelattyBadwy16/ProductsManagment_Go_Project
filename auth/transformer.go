package auth

import "project/Database/models"

func BrandTransformer(row models.User) AccountResponse {
	return AccountResponse{
		ID:   row.ID,
		Name: row.Name,
	}
}
