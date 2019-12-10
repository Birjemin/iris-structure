package main

import (
	"flag"
	"github.com/birjemin/iris-structure/command"
	"github.com/birjemin/iris-structure/conf"
	"github.com/birjemin/iris-structure/utils"
	"github.com/birjemin/iris-structure/web/routes"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"sync"
)

func main() {
	var port = flag.String("port", "8081", "port flag")
	flag.Parse()

	// set level
	golog.SetLevel(conf.Sysconfig.LoggerLevel)

	app := newApp()
	// route
	routes.InitRouter(app)

	// 启动一个协程
	// 全局退出变量
	var ConsumerExitSignal = false
	// 全局信号
	var WgConsumerExit sync.WaitGroup
	if conf.Sysconfig.ConsumerNum != 0 {
		WgConsumerExit.Add(1)
		// 启动一个常驻协程，消耗redis队列数据
		go command.Consumer(1, &WgConsumerExit, &ConsumerExitSignal)
	}

	// 启动一个常驻协程，来优雅的关闭
	go utils.GracefulShutdown(app, &WgConsumerExit, &ConsumerExitSignal)

	// Note:
	// WithoutInterruptHandler:Disable the default behavior with the option `WithoutInterruptHandler`
	// and register a new interrupt handler (globally, across all possible hosts).
	//
	// iris.WithoutServerError(iris.ErrServerClosed): WithoutServerError will cause to ignore the matched "errors"
	// from the main application's `Run` function.
	_ = app.Run(iris.Addr(":"+(*port)), iris.WithoutInterruptHandler)
}

func newApp() *iris.Application {
	// recovers on panics and logs the incoming http requests.
	app := iris.Default()

	// optimization
	app.Configure(iris.WithOptimizations)

	// cors
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: false,
		AllowedHeaders:   []string{"*"},
	})
	app.Use(crs)
	app.AllowMethods(iris.MethodOptions)
	return app
}
