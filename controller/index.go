package controller

import (
	"fmt"
	model "pornplay/models"

	"github.com/kataras/iris"
)

var category model.Category

// Indexs index data
func Indexs(ctx iris.Context) {
	data, err := category.GetIndexs()
	if err != nil {
		fmt.Println("err")
	}
	ctx.JSON(iris.Map{"category": data})
}
