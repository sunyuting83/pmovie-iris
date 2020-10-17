package main

import (
	stdCtx "context"
	"flag"
	router "porn_movie/router"
	"strings"
	"sync"
	"time"

	"github.com/kataras/iris/v12"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "3000", "端口号，默认为3000")
	flag.Parse()
	app := router.InitRouter()
	// 优雅的关闭程序
	serverWG := new(sync.WaitGroup)
	defer serverWG.Wait()

	iris.RegisterOnInterrupt(func() {
		serverWG.Add(1)
		defer serverWG.Done()

		ctx, cancel := stdCtx.WithTimeout(stdCtx.Background(), 20*time.Second)
		defer cancel()

		// 关闭所有主机
		app.Shutdown(ctx)
	})

	// app.Run(iris.Addr(strings.Join([]string{":", port}, "")), iris.WithoutServerError(iris.ErrServerClosed))
	app.Listen(strings.Join([]string{":", port}, ""))
}
