package models

import "time"

type Product struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `json:"name"`
	BrandID uint   `json:"brand_id"`
	Brand   Brand  `gorm:"foreignKey:BrandID;references:ID"`
	Tags    []Tag  `gorm:"many2many:product_tags;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
