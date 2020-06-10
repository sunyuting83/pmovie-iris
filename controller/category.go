package controller

import (
	"encoding/json"
	"fmt"
	model "pornplay/models"
	"strings"

	"github.com/kataras/iris"
)

var categoryList model.CategoryList

// Category Category List
type Category struct {
	ID     int64  `json:"id"`
	CID    int64  `json:"cid"`
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	More   More   `json:"more"`
	Region string `json:"region"`
}

// More more
type More struct {
	Play string   `json:"play"`
	Lg   string   `json:"lg"`
	CIM  string   `json:"cim"`
	Text []string `json:"text"`
}

// CategoryList Category List
func CategoryList(ctx iris.Context) {
	// fmt.Println(ctx.Host())
	category, cerr := ctx.URLParamInt64("category")
	if cerr != nil {
		category = 1
	}
	page, perr := ctx.URLParamInt64("page")
	if perr != nil {
		page = 1
	}
	data, err := categoryList.GetCategory(category, page)
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
		if !strings.Contains(more.CIM, "http") {
			more.CIM = strings.Join([]string{"https", more.CIM}, ":")
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
	ctx.JSON(iris.Map{"movie": d})
}

// CategoryTojson fun
func CategoryTojson(s string) (p More) {
	if err := json.Unmarshal([]byte(s), &p); err != nil {
		// fmt.Println(err.Error())
		return
	}
	return
}
