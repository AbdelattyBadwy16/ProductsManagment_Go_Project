package tag

import (
	"project/Database"
	"project/Database/models"
	"project/common"

	"gorm.io/gorm"
)

func storeTagRow(row models.Tag) (models.Tag, error) {
	db := Database.DB //need to call db var here
	err := db.Save(&row).Error
	return row, err
}

func AllTags() ([]models.Tag, error) {
	db := Database.DB
	var brands []models.Tag
	err := db.Find(&brands).Error
	return brands, err
}

func DeleteTagRow(request DeleteRequest) error {
	db := Database.DB
	var tag models.Tag
	err := db.First(&tag, request.Id).Error
	if err != nil {
		return err
	}
	err = db.Delete(&tag).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateTagRow(row models.Tag) (models.Tag, error) {
	db := Database.DB
	err := db.Save(&row).Error
	return row, err
}

func Filter(request FilterRequest) ([]models.Tag, error) {
	db := Database.DB
	if request.Search != "" {
		db = db.Where("name LIKE ?", "%"+request.Search+"%")
	}
	db = filterByDate(request.StartDate, request.EndDate, db)
	if request.Column != "" && request.Type != "" {
		db = common.OrderBy(request.Column, request.Type, db)
	}
	var rows []models.Tag
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
