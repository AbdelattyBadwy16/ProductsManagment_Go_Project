package tag

import "github.com/kataras/iris/v12"

func Routes(app *iris.Application) {
	group := app.Party("/tag")
	group.Get("/all", GetAllTags)
	group.Post("/create", CreateTag)
	group.Patch("/update",UpdateTag)
	group.Delete("/delete",DeleteTag)
	group.Get("/show",FilterTag)
}
