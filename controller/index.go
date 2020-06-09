package controller

import (
	"fmt"
	model "pornplay/models"
	"strings"

	"github.com/kataras/iris"
)

var category model.Category

// Indexs index data
func Indexs(ctx iris.Context) {
	// fmt.Println(ctx.Host())
	var datalen int
	data, err := category.GetIndexs()
	if len(data) <= 0 {
		datalen = 0
	} else {
		datalen = len(data)
	}
	d := make([]model.Category, datalen)
	if err != nil {
		fmt.Println("err")
	}
	host := strings.Join([]string{"http", ctx.Host()}, "://")
	for _, item := range data {
		item.Cover = strings.Join([]string{host, item.Cover}, "")
		d = append(d, model.Category{ID: item.ID, Category: item.Category, Sort: item.Sort, Cover: item.Cover})
	}
	ctx.JSON(iris.Map{"category": d})
}
