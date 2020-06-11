package router

import (
	controller "pornplay/controller"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

// InitRouter make router
func InitRouter() *iris.Application {
	app := iris.New()
	// app := iris.Default()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.HandleDir("/cover", "./static/cover") // static
	// Simple group: v1.
	v1 := app.Party("/api", crs).AllowMethods(iris.MethodOptions) // <- 对于预检很重要。
	{
		v1.Get("/", controller.Indexs)
		v1.Get("/category", controller.CategoryList)
		v1.Get("/searchkey", controller.SearchKey)
		v1.Get("/search", controller.GetSearch)
		v1.Get("/hotkey", controller.GetSearchHot)
		v1.Get("/gelivels", controller.GetLives)
		v1.Get("/gelive", controller.GetLive)
	}
	return app
}
