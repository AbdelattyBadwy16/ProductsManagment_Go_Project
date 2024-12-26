package product

import "project/Database/models"

func ProductTransformer(row models.Product) ProductResponse {
	return ProductResponse{
		ID:         row.ID,
		Name:       row.Name,
		BrandID:    row.BrandID,
		Created_At: row.CreatedAt,
	}
}

func ProductsTransformer(rows []models.Product) []ProductResponse {
	Product := make([]ProductResponse, 0)
	for _, row := range rows {
		Product = append(Product, ProductTransformer(row))
	}
	return Product
}
