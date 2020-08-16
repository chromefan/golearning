package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/panjf2000/ants"
)

var tunnel = make(chan string, 1)

type workerPool struct {
	// Function for serving server connections.
	// It must leave c unclosed.
	WorkerFunc ServeHandler

	MaxWorkersCount int

	LogAllErrors bool

	MaxIdleWorkerDuration time.Duration

	Logger Logger

	lock         sync.Mutex
	workersCount int
	mustStop     bool

	ready []*workerChan

	stopCh chan struct{}

	workerChanPool sync.Pool

	connState func(net.Conn, ConnState)
}
func main() {
	go IncomingZombie()

	chairPool, _ := ants.NewPoolWithFunc(3, ExecuteZombie) // 声明有几把电椅
	defer chairPool.Release()

	for {
		select {
		case a := <-tunnel:
			go chairPool.Invoke(a)
		}
	}
}

// 处决僵尸
func ExecuteZombie(i interface{}) {
	fmt.Printf("正在处决僵尸 %s 号，还有5秒钟....\n", i.(string))
	time.Sleep(5*time.Second)
	fmt.Printf(":) %s 玩完了，下一个\n-----------------\n", i.(string))
}

// 僵尸不断进来
func IncomingZombie() {
	for i := 0; i < 10; i++ {
		tunnel <- strconv.Itoa(i)
	}
}