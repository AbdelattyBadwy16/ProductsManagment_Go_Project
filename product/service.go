package product

import (
	"encoding/json"
	"project/Database/models"
	redis "project/ExternalService/Redis"
)

func GetAllProductsService() ([]models.Product, error) {

	// check cache
	res, err := redis.GetValue("products")
	if err == nil {
		var products []models.Product
		err = json.Unmarshal([]byte(res), &products)
		return products, nil
	}

	// cache not found
	products, err := AllProduct()
	if err != nil {
		return []models.Product{}, err
	}

	redis.SetValue("Products", products)
	return products, nil
}

func CreateProductService(request CreateRequest) (models.Product, error) {
	product := setProduct(request)
	product, err := setProductRow(product)
	if err != nil {
		return models.Product{}, err
	}
	// Clear cache
	redis.RemoveValue("Products")
	return product, nil
}

func UpdateProductService(request UpdateRequest) (models.Product, error) {
	product := setProductUpdate(request)
	product, err := updateProductRow(product)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func DeleteProductService(request DeleteRequest) (models.Product, error) {
	product, err := deleteProductRow(request)
	if err != nil {
		return models.Product{}, err
	}
	// Clear cache
	redis.RemoveValue("Products")
	return product, nil
}

func FilterProductService(request FilterRequest) ([]models.Product, error) {
	product, err := Filter(request)
	if err != nil {
		return []models.Product{}, err
	}
	return product, nil
}
