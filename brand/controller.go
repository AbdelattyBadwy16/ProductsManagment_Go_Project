package brand

import (
	"github.com/kataras/iris/v12"
)

func CreateBrand(c iris.Context) {
	var request CreateRequest
	if err := c.ReadJSON(&request); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	err := ValidateCreateRequest(request)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	brand, err := CreateBrandService(request)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	response := BrandTransformer(brand)
	c.StatusCode(iris.StatusCreated)
	c.JSON(response)
}

func GetAllBrands(c iris.Context) {
	brands, err := GetALlBrandsService()
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	response := BrandsTransformer(brands)
	c.StatusCode(iris.StatusCreated)
	c.JSON(response)
}

func DeleteBrand(c iris.Context) {
	var request DeleteRequest
	if err := c.ReadJSON(&request); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	err := ValidateDeleteRequest(request)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	err = DeleteBrandService(request)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	c.StatusCode(iris.StatusCreated)
}

func UpdateBrand(c iris.Context) {
	var request UpdateRequest
	if err := c.ReadJSON(&request); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	err := ValidateUpdateRequest(request)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	brand, err := UpdateBrandService(request)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	c.StatusCode(iris.StatusCreated)
	c.JSON(brand)
}

func FilterBrand(c iris.Context) {
	var request FilterRequest
	if err := c.ReadJSON(&request); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}

	products, err := FilterBrandService(request)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	c.StatusCode(iris.StatusCreated)
	c.JSON(products)
}
