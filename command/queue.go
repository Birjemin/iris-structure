package command

import (
	"fmt"
	"github.com/birjemin/iris-structure/cache"
	"github.com/kataras/golog"
	"math/rand"
	"sync"
	"time"
)

const QueueName string = "xxxx"

func Consumer(num int, wgc *sync.WaitGroup, exitSignal *bool) {
	rand.Seed(time.Now().UnixNano())
	// 消费者子协程同步信号
	for i := 0; i < num; i++ {
		go func(wgChild *sync.WaitGroup, i int) {
			// 每个协程死循环消费队列数据，有终止信号即跳出循环
			for {
				data, err := cache.LPop(QueueName)
				if err != nil {
					time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
					fmt.Println("consume element:", i)
				} else {
					fmt.Println("consume element:", data)
				}
				if *exitSignal {
					golog.Error("[queue]Consume exited, No is : ", i)
					break
				}
			}
			wgChild.Done()
		}(wgc, i)
	}
}
