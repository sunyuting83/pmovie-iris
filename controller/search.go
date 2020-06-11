package controller

import (
	"fmt"
	"strings"

	"github.com/kataras/iris"
)

// SearchKey Search Key
func SearchKey(ctx iris.Context) {
	// fmt.Println(ctx.Host())
	key := ctx.URLParamTrim("key")

	data, err := categoryList.SearchKey(key)
	if err != nil {
		fmt.Println("err")
	}
	d := make([]Category, 0)
	for _, item := range data {
		more := CategoryTojson(item.More)
		// fmt.Println(more, item.More)
		if !strings.Contains(item.Cover, "http") {
			item.Cover = strings.Join([]string{"https", item.Cover}, ":")
		}
		if len(more.CIM) > 0 {
			if !strings.Contains(more.CIM, "http") {
				more.CIM = strings.Join([]string{"https", more.CIM}, ":")
			}
		}
		if len(more.Play) > 0 {
			if !strings.Contains(more.Play, "http") {
				more.Play = strings.Join([]string{"http", more.Play}, ":")
			}
		}
		d = append(d, Category{
			ID:     item.ID,
			CID:    item.CID,
			Title:  item.Title,
			Cover:  item.Cover,
			Region: item.Region,
			More:   more,
		})
	}
	ctx.JSON(d)
}
