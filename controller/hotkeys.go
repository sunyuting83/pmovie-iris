package controller

import (
	"encoding/json"
	leveldb "pornplay/leveldb"
	"sort"

	"github.com/kataras/iris"
)

// HotKey hot key
type HotKey struct {
	Key   string `json:"key"`
	Click int64  `json:"click"`
}

// HotKeys s
type HotKeys []HotKey

//Len()
func (s HotKeys) Len() int {
	return len(s)
}

// Less 排序
func (s HotKeys) Less(i, j int) bool {
	return s[i].Click > s[j].Click
}

//Swap()
func (s HotKeys) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// GetSearchHot fun
func GetSearchHot(ctx iris.Context) {
	hotkey := GetHotKey()
	ctx.JSON(hotkey)
}

// GetHotKey (hotkey map[]in)
func GetHotKey() []string {
	var (
		hotkey HotKeys
		list   []string
	)
	cache := leveldb.GetLevel("hotkey")
	if cache == "leveldb: not found" {
		return list
	}
	hotkey = HotToJsons(cache)
	if len(hotkey) >= 10 {
		hotkey = hotkey[0:10]
	}
	for _, item := range hotkey {
		list = append(list, item.Key)
	}
	return list
}

// HotToJsons fun
func HotToJsons(s string) HotKeys {
	var result HotKeys
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return result
	}
	return result
}

// HotToString fun
func HotToString(d HotKeys) (result string) {
	resultByte, errError := json.Marshal(d)
	result = string(resultByte)
	if errError != nil {
		return result
	}
	return result
}

// SaveHotKey fun
func SaveHotKey(key string) {
	var (
		hotkey HotKeys
	)
	cache := leveldb.GetLevel("hotkey")
	if cache == "leveldb: not found" {
		hotkey = append(hotkey, HotKey{
			Key:   key,
			Click: 1,
		})
		leveldb.SetLevel("hotkey", HotToString(hotkey), -1)
	} else {
		hotkey = HotToJsons(cache)
		h := ChecKey(hotkey, key)
		if h {
			for i, item := range hotkey {
				if item.Key == key {
					// fmt.Println(hotkey[i])
					hotkey[i].Click = item.Click + 1
					break
				}
			}
		} else {
			hotkey = append(hotkey, HotKey{
				Key:   key,
				Click: 1,
			})
		}
		sort.Sort(hotkey)
		leveldb.SetLevel("hotkey", HotToString(hotkey), -1)
	}
	return
}

// ChecKey check key
func ChecKey(arr HotKeys, k string) bool {
	for _, item := range arr {
		if item.Key == k {
			return true
		}
	}
	return false
}
