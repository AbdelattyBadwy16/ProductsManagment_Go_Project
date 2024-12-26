package brand

import "time"

type BrandResponse struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Created_At time.Time `json:"created_at"`
}
