package product

import "project/Database/models"

func setProduct(request CreateRequest) models.Product {
	return models.Product{
		Name:    request.Name,
		BrandID: request.BrandID,
	}
}

func setProductUpdate(request UpdateRequest) models.Product {
	return models.Product{
		ID:      uint(request.Id),
		Name:    request.Name,
		BrandID: request.BrandID,
	}
}
