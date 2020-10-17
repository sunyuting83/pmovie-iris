package controller

import (
	"encoding/json"
	leveldb "porn_movie/leveldb"
	model "porn_movie/models"
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
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
	cat := strconv.FormatInt(category, 10)
	pa := strconv.FormatInt(page, 10)
	d := make([]Category, 0)
	cache := leveldb.GetLevel(strings.Join([]string{"category", cat, pa}, ":"))
	if cache == "leveldb: not found" {
		data, err := categoryList.GetCategory(category, page)
		if err != nil {
			d = make([]Category, 0)
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
		if len(data) > 0 {
			leveldb.SetLevel(strings.Join([]string{"category", cat, pa}, ":"), CateToStr(d), 86400000)
		}
	} else {
		d = CateToJsons(cache)
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

// CateToJsons json
func CateToJsons(s string) (result []Category) {
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return
	}
	return
}

// CateToStr json
func CateToStr(d []Category) (result string) {
	resultByte, errError := json.Marshal(d)
	result = string(resultByte)
	if errError != nil {
		return
	}
	return
}
