package utils

import (
	"context"
	"github.com/birjemin/iris-structure/datasource"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// 消费者父协程退出信号
//var WgConsumerExit sync.WaitGroup
// 消费者子协程退出信号
//var QueueExitSignal = false

func GracefulShutdown(app *iris.Application, wgc *sync.WaitGroup, exitSignal *bool) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch,
		// kill -SIGINT XXXX or Ctrl+c
		os.Interrupt,
		syscall.SIGINT, // register that too, it should be ok
		// os.Kill  is equivalent with the syscall.Kill
		os.Kill,
		syscall.SIGKILL, // register that too, it should be ok
		// kill -SIGTERM XXXX
		syscall.SIGTERM,
	)
	select {
	case <-ch:
		timeout := 5 * time.Second

		// 通知消费者协程退出~~~
		*exitSignal = true
		// 等待消费者协程信号结束（阻塞）
		wgc.Wait()

		var err error
		// Close database, redis, truncate message queues, etc.
		err = datasource.CloseDb()
		golog.Error("[main]DB Pool Exited...", err)
		err = datasource.CloseRedis()
		golog.Error("[main]Redis Pool Exited...", err)
		err = datasource.CloseGRPC()
		golog.Error("[main]GRPC Exited...", err)
		golog.Error("[main]Iris shutdown...")
		golog.Error("[main]Defer, Pool Redis and Db Stats", datasource.StatsRedis(), datasource.StatsDB())

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		_ = app.Shutdown(ctx)
	}
}
