package models

import "time"

type Tag struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:product_tags;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
