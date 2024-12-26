package tag

import "project/Database/models"

func setTag(request CreateRequest) models.Tag {
	return models.Tag{
		Name: request.Name,
	}
}

func setTagUpdate(request UpdateRequest) models.Tag {
	return models.Tag{
		ID: uint(request.Id),
		Name: request.Name,
	}
}
