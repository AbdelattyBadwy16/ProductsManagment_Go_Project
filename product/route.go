package product

import "github.com/kataras/iris/v12"

func Routes(app *iris.Application) {
	group := app.Party("/product")
	group.Get("/all", GetAllProducts)
	group.Post("/create", CreateProduct)
	group.Patch("/update", UpdateProduct)
	group.Delete("/delete", DeleteProduct)
	group.Get("/show",FilterProduct)
}
