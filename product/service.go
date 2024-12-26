package product

import (
	"project/Database/models"
)

func GetAllProductsService() ([]models.Product, error) {
	products, err := AllProduct()
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

func CreateProductService(request CreateRequest) (models.Product, error) {
	product := setProduct(request)
	product, err := setProductRow(product)
	if err != nil {
		return models.Product{}, err
	}
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
	return product, nil
}

func FilterProductService(request FilterRequest) ([]models.Product, error) {
	product, err := Filter(request)
	if err != nil {
		return []models.Product{}, err
	}
	return product, nil
}

