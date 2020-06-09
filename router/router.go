package router

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

// InitRouter make router
func InitRouter() *iris.Application {
	app := iris.New()
	// app := iris.Default()

	app.Use(recover.New())
	app.Use(logger.New())
	// Simple group: v1.
	v1 := app.Party("/api")
	{
		v1.Get("/login", func(ctx iris.Context) {
			ctx.JSON(iris.Map{"message": "Hello Iris!"})
		})
		v1.Get("/submit", func(ctx iris.Context) {
			ctx.JSON(iris.Map{"message": "Hello Iris!"})
		})
		v1.Get("/read", func(ctx iris.Context) {
			ctx.JSON(iris.Map{"message": "Hello Iris!"})
		})
	}
	return app
}
