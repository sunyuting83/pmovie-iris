package controller

import (
	"encoding/json"
	"fmt"
	leveldb "pornplay/leveldb"
	model "pornplay/models"

	"github.com/kataras/iris"
)

var category model.Category

// Indexs index data
func Indexs(ctx iris.Context) {
	// fmt.Println(ctx.Host())
	d := make([]model.Category, 0)
	cache := leveldb.GetLevel("index")
	if cache == "leveldb: not found" {
		data, err := category.GetIndexs()
		if err != nil {
			fmt.Println("err")
		}
		d = data
		// host := strings.Join([]string{"http", ctx.Host()}, "://")
		// for _, item := range data {
		// 	item.Cover = strings.Join([]string{host, item.Cover}, "")
		// 	d = append(d, model.Category{ID: item.ID, Category: item.Category, Sort: item.Sort, Cover: item.Cover})
		// }
		if len(data) > 0 {
			leveldb.SetLevel("index", InToStr(d), 86400000)
		}
	} else {
		d = InToJsons(cache)
	}
	ctx.JSON(iris.Map{"category": d})
}

// InToJsons index str to json
func InToJsons(s string) (result []model.Category) {
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return
	}
	return
}

// InToStr fun
func InToStr(d []model.Category) (result string) {
	resultByte, errError := json.Marshal(d)
	result = string(resultByte)
	if errError != nil {
		return
	}
	return
}
