package brand

import (
	"github.com/kataras/iris/v12"
)

func Routes(app *iris.Application) {
	group := app.Party("/brand")
	group.Get("/all", GetAllBrands)
	group.Post("/create", CreateBrand)
	group.Delete("/delete", DeleteBrand)
	group.Patch("/update", UpdateBrand)
	group.Get("/show",FilterBrand)
}
