package brand

import "project/Database/models"

func BrandTransformer(row models.Brand) BrandResponse {
	return BrandResponse{
		ID:   row.ID,
		Name: row.Name,
		Created_At : row.CreatedAt,
	}
}

func BrandsTransformer(rows []models.Brand) []BrandResponse {
	brands := make([]BrandResponse, 0)
	for _, row := range rows {
		brands = append(brands, BrandTransformer(row))
	}
	return brands
}
