package auth

import (
	"github.com/kataras/iris/v12"
)

func Routes(app *iris.Application) {
	group := app.Party("/auth")
	group.Post("/register", CreateAccount)
	group.Post("/login",Login)
}
