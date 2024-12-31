package brand

import (
	"encoding/json"
	"project/Database/models"
	redis "project/ExternalService/Redis"
)

func CreateBrandService(request CreateRequest) (models.Brand, error) {
	brand := setBrand(request)
	brand, err := storeBrandRow(brand)
	if err != nil {
		return models.Brand{}, err
	}
	// Clear cache
	redis.RemoveValue("Brands")
	return brand, nil
}

func GetALlBrandsService() ([]models.Brand, error) {
	// check cache
	res, err := redis.GetValue("Brands")
	if err == nil {
		var brands []models.Brand
		err = json.Unmarshal([]byte(res), &brands)
		return brands, nil
	}
	brands, err := AllBrand()
	if err != nil {
		return []models.Brand{}, err
	}
	redis.SetValue("Brandsg", brands)
	return brands, nil
}

func DeleteBrandService(request DeleteRequest) error {
	err := DeleteBrandRow(request)
	if err != nil {
		return err
	}
	// Clear cache
	redis.RemoveValue("Brands")
	return nil
}

func UpdateBrandService(request UpdateRequest) (models.Brand, error) {
	brand := setBrandUpdate(request)
	brand, err := UpdateBrandRow(brand)
	if err != nil {
		return models.Brand{}, err
	}
	return brand, nil
}

func FilterBrandService(request FilterRequest) ([]models.Brand, error) {
	brand, err := Filter(request)
	if err != nil {
		return []models.Brand{}, err
	}
	return brand, nil
}
