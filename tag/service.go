package tag

import (
	"project/Database/models"
)

func CreateTagService(request CreateRequest) (models.Tag, error) {
	tag := setTag(request)
	tag, err := storeTagRow(tag)
	if err != nil {
		return models.Tag{}, err
	}
	return tag, nil
}

func GetAllTagsService() ([]models.Tag, error) {
	tags, err := AllTags()
	if err != nil {
		return []models.Tag{}, err
	}
	return tags, nil
}

func DeleteTagService(request DeleteRequest) error {
	err := DeleteTagRow(request)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTagService(request UpdateRequest) (models.Tag, error) {
	tag := setTagUpdate(request)
	tag, err := UpdateTagRow(tag)
	if err != nil {
		return models.Tag{}, err
	}
	return tag, nil
}

func FilterTagService(request FilterRequest) ([]models.Tag, error) {
	tag, err := Filter(request)
	if err != nil {
		return []models.Tag{}, err
	}
	return tag, nil
}

