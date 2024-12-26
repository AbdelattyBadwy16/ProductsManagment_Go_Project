package brand

import "project/Database/models"

func setBrand(request CreateRequest) models.Brand {
	return models.Brand{
		Name: request.Name,
	}
}

func setBrandUpdate(request UpdateRequest) models.Brand {
	return models.Brand{
		ID: uint(request.Id),
		Name: request.Name,
	}
}
