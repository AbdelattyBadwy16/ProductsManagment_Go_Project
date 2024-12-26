package product

import (
	"project/Database"
	"project/Database/models"
	"project/common"

	"gorm.io/gorm"
)

func AllProduct() ([]models.Product, error) {
	db := Database.DB
	var products []models.Product
	err := db.Preload("Brand").Find(&products).Error
	return products, err
}

func setProductRow(row models.Product) (models.Product, error) {
	db := Database.DB
	var brand models.Brand
	err := db.First(&brand, row.BrandID).Error
	if err != nil {
		return models.Product{}, err
	}
	row.Brand = brand
	err = db.Save(&row).Error
	return row, err
}

func updateProductRow(row models.Product) (models.Product, error) {
	db := Database.DB
	var brand models.Brand
	err := db.First(&brand, row.BrandID).Error
	if err != nil {
		return models.Product{}, err
	}
	row.Brand = brand
	err = db.Save(&row).Error
	return row, err
}

func deleteProductRow(request DeleteRequest) (models.Product, error) {
	db := Database.DB
	var product models.Product
	err := db.First(&product, request.Id).Error
	if err != nil {
		return models.Product{}, err
	}
	err = db.Delete(&product).Error
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}


func Filter(request FilterRequest) ([]models.Product, error) {
	db := Database.DB
	if request.Search != "" {
		db = db.Where("name LIKE ?", "%"+request.Search+"%")
	}
	db = filterByDate(request.StartDate, request.EndDate, db)
	if request.Column != "" && request.Type != "" {
		db = common.OrderBy(request.Column, request.Type, db)
	}
	var rows []models.Product
	err := db.Find(&rows).Error
	return rows, err
}

func filterByDate(startDate, endDate string, db *gorm.DB) *gorm.DB {
	if startDate != "" {
		db = common.FilterByField("created_at", ">=", startDate, db)
	}
	if startDate != "" {
		db = common.FilterByField("created_at", "<=", endDate, db)
	}
	return db
}
