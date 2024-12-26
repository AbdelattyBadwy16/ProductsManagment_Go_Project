package tag

type CreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type DeleteRequest struct {
	Id int `json:"id" validate:"required"`
}

type UpdateRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type FilterRequest struct {
	Search    string `json:"search"`
	StartDate string `json:"start_date" example:"2023-10-20T00:00:00Z"`
	EndDate   string `json:"end_date" example:"2023-10-20T22:00:00Z"`
	Column    string `json:"column"`
	Type      string `json:"type"`
}
