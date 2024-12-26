package tag

import "project/Database/models"

func TagTransformer(row models.Tag) TagResponse {
	return TagResponse{
		ID:   row.ID,
		Name: row.Name,
	}
}

func TagsTransformer(rows []models.Tag) []TagResponse {
	brands := make([]TagResponse, 0)
	for _, row := range rows {
		brands = append(brands, TagTransformer(row))
	}
	return brands
}
