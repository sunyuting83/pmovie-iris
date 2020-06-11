package controller

import (
	"encoding/json"
	"net/http"
	leveldb "pornplay/leveldb"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/kataras/iris"
)

// LiveList Live List
type LiveList struct {
	URL   string `json:"url"`
	Title string `josn:"title"`
	Cover string `json:"cover"`
}

// Live live
type Live struct {
	URL  string `json:"url"`
	Code int    `json:"code"`
}

// GetLives Get Live List
func GetLives(ctx iris.Context) {
	cache := leveldb.GetLevel("live")
	lives := make([]LiveList, 0)
	if cache == "leveldb: not found" {
		lives = LiveScrape(false)
	} else {
		lives = LiveToJsons(cache)
	}
	ctx.JSON(lives)
}

// GetLive Get Live List
func GetLive(ctx iris.Context) {
	t := ctx.URLParamTrim("t")
	live := getLiveScrape(t, false)
	ctx.JSON(live)
}

// LiveScrape get live list
func LiveScrape(cors bool) []LiveList {
	// Request the HTML page.
	var err error
	var url string
	List := make([]LiveList, 0)
	url = "https://members.sexcamvideos.net"
	if cors {
		url = strings.Join([]string{"https://cors.zme.ink", url}, "/")
	}
	// fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return List
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return List
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return List
	}
	root := doc.Find("body ul#room_list > li.room_list_room")
	imgurl := "https://roomimg.stream.highwebmedia.com/ri"
	root.Each(func(index int, ele *goquery.Selection) {
		title := ele.Find("div.details > div.title > a").Text()
		title = strings.Replace(title, " ", "", -1)
		t := time.Now().Unix()
		tt := strconv.FormatInt(t, 10)
		cc := strings.Join([]string{imgurl, title}, "/")
		cc = strings.Join([]string{cc, "jpg?"}, ".")
		cover := strings.Join([]string{cc, tt}, "")
		urls := strings.Join([]string{url, title}, "/")
		if index <= 16 {
			List = append(List, LiveList{URL: urls, Title: title, Cover: cover})
		}
	})
	if len(List) > 0 {
		leveldb.SetLevel("live", LiveToStr(List), 300000)
	}
	return List
}

// LiveScrape get live list
func getLiveScrape(t string, cors bool) Live {
	// Request the HTML page.
	var (
		err  error
		url  string
		live Live
	)
	live = Live{URL: "", Code: 0}
	if len(t) <= 0 {
		return live
	}
	url = strings.Join([]string{"https://members.sexcamvideos.net", t}, "/")
	if cors {
		url = strings.Join([]string{"https://cors.zme.ink", url}, "/")
	}
	// fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		return live
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return live
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return live
	}
	doc.Find("body script").Each(func(index int, ele *goquery.Selection) {
		if strings.Contains(ele.Text(), "window.initialRoomDossier") {
			s := ele.Text()
			x := strings.Index(s, `hls_source\u0022: \u0022`) + 24
			y := strings.Index(s, `\u0022, \u0022allow_show_recordings`)
			s = s[x:y]
			// s := ele.Text()
			s = strings.Replace(s, `\u002D`, `-`, -1)
			live = Live{URL: s, Code: 1}
		}
	})

	return live
}

// LiveToJsons live to json
func LiveToJsons(s string) (result []LiveList) {
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return
	}
	return
}

// LiveToStr fun
func LiveToStr(d []LiveList) (result string) {
	resultByte, errError := json.Marshal(d)
	result = string(resultByte)
	if errError != nil {
		return
	}
	return
}
