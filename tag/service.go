package tag

import (
	"encoding/json"
	"project/Database/models"
	redis "project/ExternalService/Redis"
)

func CreateTagService(request CreateRequest) (models.Tag, error) {
	tag := setTag(request)
	tag, err := storeTagRow(tag)
	if err != nil {
		return models.Tag{}, err
	}
	// Clear cache
	redis.RemoveValue("Tags")
	return tag, nil
}

func GetAllTagsService() ([]models.Tag, error) {
	// check cache
	res, err := redis.GetValue("Tags")
	if err == nil {
		var tags []models.Tag
		err = json.Unmarshal([]byte(res), &tags)
		return tags, nil
	}
	tags, err := AllTags()
	if err != nil {
		return []models.Tag{}, err
	}
	redis.SetValue("Tags", tags)
	return tags, nil
}

func DeleteTagService(request DeleteRequest) error {
	err := DeleteTagRow(request)
	if err != nil {
		return err
	}
	// Clear cache
	redis.RemoveValue("Tags")
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

