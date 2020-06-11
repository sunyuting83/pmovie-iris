package controller

import (
	"fmt"
	leveldb "pornplay/leveldb"
	"strings"

	"github.com/kataras/iris"
)

// SearchKey Search Key
func SearchKey(ctx iris.Context) {
	// fmt.Println(ctx.Host())
	key := ctx.URLParamTrim("key")

	d := make([]Category, 0)
	cache := leveldb.GetLevel(strings.Join([]string{"searchkey", key}, ":"))
	if cache == "leveldb: not found" {
		data, err := categoryList.SearchKey(key)
		if err != nil {
			fmt.Println("err")
		}
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
		leveldb.SetLevel(strings.Join([]string{"searchkey", key}, ":"), CateToStr(d), 86400000)
	} else {
		d = CateToJsons(cache)
	}
	ctx.JSON(d)
}

// GetSearch Search List
func GetSearch(ctx iris.Context) {
	// fmt.Println(ctx.Host())
	key := ctx.URLParamTrim("key")
	page, perr := ctx.URLParamInt64("page")
	if perr != nil {
		page = 1
	}

	data, err := categoryList.Search(key, page)
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
		if len(d) > 0 {
			SaveHotKey(key)
		}
	}
	ctx.JSON(d)
}
