package common

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

func ValidateRequest(request interface{}) error {
	err := validate.Struct(request)
	if err != nil {
		return err
	}
	return nil
}

func FilterByField(field string, operator string, value interface{}, db *gorm.DB) *gorm.DB {
	return db.Where(field+" "+operator+" ?", value)
}

func OrderBy(Column string, Type string, db *gorm.DB) *gorm.DB {
	db = db.Order(Column + " " + Type)
	return db
}
