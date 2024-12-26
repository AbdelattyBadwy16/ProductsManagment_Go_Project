package brand

import (
	"project/Database/models"
)

func CreateBrandService(request CreateRequest) (models.Brand, error) {
	brand := setBrand(request)
	brand, err := storeBrandRow(brand)
	if err != nil {
		return models.Brand{}, err
	}
	return brand, nil
}

func GetALlBrandsService() ([]models.Brand, error) {
	brand, err := AllBrand()
	if err != nil {
		return []models.Brand{}, err
	}
	return brand, nil
}

func DeleteBrandService(request DeleteRequest) error {
	err := DeleteBrandRow(request)
	if err != nil {
		return err
	}
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
