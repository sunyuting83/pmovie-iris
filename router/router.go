package router

import (
	controller "pornplay/controller"

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
	app.HandleDir("/cover", "./static/cover") // static
	// Simple group: v1.
	v1 := app.Party("/api")
	{
		v1.Get("/", controller.Indexs)
	}
	return app
}
