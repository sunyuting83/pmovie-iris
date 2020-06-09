package controller

import "github.com/kataras/iris"

// Indexs index data
func Indexs(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "Hello Iris!"})
}
