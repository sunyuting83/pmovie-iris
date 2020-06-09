package main

import (
	"flag"
	router "pornplay/router"
	"strings"

	"github.com/kataras/iris"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "3000", "端口号，默认为3000")
	flag.Parse()
	app := router.InitRouter()
	app.Run(iris.Addr(strings.Join([]string{":", port}, "")), iris.WithoutServerError(iris.ErrServerClosed))
}
