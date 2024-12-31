package main

import (
	"project/Database"
	redis "project/ExternalService/Redis"
	"project/brand"
	"project/product"
	"project/tag"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	godotenv.Load()
	Database.InitDB()
	redis.Init()
	app := iris.New()
	// Custom CORS middleware
	app.UseRouter(func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if ctx.Method() == "OPTIONS" {
			ctx.StatusCode(204)
			return
		}
		ctx.Next()
	})
	product.Routes(app)
	brand.Routes(app)
	tag.Routes(app)

	app.Listen(":8081")
}
