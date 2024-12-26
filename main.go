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

	product.Routes(app)
	brand.Routes(app)
	tag.Routes(app)

	app.Listen(":8081")
}
