package tag

import (
	"github.com/kataras/iris/v12"
)

func CreateTag(c iris.Context) {
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
	tag, err := CreateTagService(request)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	response := TagTransformer(tag)
	c.StatusCode(iris.StatusCreated)
	c.JSON(response)
}

func GetAllTags(c iris.Context) {
	tags, err := GetAllTagsService()
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	response := TagsTransformer(tags)
	c.StatusCode(iris.StatusCreated)
	c.JSON(response)
}

func DeleteTag(c iris.Context) {
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
	err = DeleteTagService(request)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	c.StatusCode(iris.StatusCreated)
}

func UpdateTag(c iris.Context) {
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
	brand, err := UpdateTagService(request)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	c.StatusCode(iris.StatusCreated)
	c.JSON(brand)
}

func FilterTag(c iris.Context) {
	var request FilterRequest
	if err := c.ReadJSON(&request); err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}

	products, err := FilterTagService(request)
	if err != nil {
		c.StatusCode(iris.StatusBadRequest)
		return
	}
	c.StatusCode(iris.StatusCreated)
	c.JSON(products)
}
