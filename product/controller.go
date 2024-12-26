package product

import (
	"github.com/kataras/iris/v12"
)

func GetAllProducts(c iris.Context) {
	products, err := GetAllProductsService()
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	response := ProductsTransformer(products)
	c.StatusCode(iris.StatusCreated)
	c.JSON(response)
}

func CreateProduct(c iris.Context) {
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
	product, err := CreateProductService(request)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	response := ProductTransformer(product)
	c.StatusCode(iris.StatusCreated)
	c.JSON(response)
}

func UpdateProduct(c iris.Context) {
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
	product, err := UpdateProductService(request)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	response := ProductTransformer(product)
	c.StatusCode(iris.StatusCreated)
	c.JSON(response)
}

func DeleteProduct(c iris.Context) {
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
	product, err := DeleteProductService(request)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	response := ProductTransformer(product)
	c.StatusCode(iris.StatusCreated)
	c.JSON(response)
}

func FilterProduct(c iris.Context) {
	var request FilterRequest
	if err := c.ReadJSON(&request); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}

	products, err := FilterProductService(request)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	c.StatusCode(iris.StatusCreated)
	c.JSON(products)
}
