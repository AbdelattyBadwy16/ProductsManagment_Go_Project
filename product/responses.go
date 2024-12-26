package product

import (
	"project/Database/models"
	"time"
)

type ProductResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	BrandID uint   `json:"brand_id"`
	Brand   models.Brand
	Created_At time.Time `json:"created_at"` 
}
