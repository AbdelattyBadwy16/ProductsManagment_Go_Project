package models

import "time"


type Brand struct {
	ID       uint `gorm:"primaryKey"`
	Name     string 
	Products []Product `gorm:"foreignKey:BrandID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

