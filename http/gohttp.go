package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	ThreadNum = 2
	SleepTime = 60
	RedisHost = "127.0.0.1:6379"
	QueueName = "user-uid-queue"
)

var wg sync.WaitGroup

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	str := values.Get("uid")
	uid, _ := strconv.ParseInt(str, 0, 64)
	rc := connRedis()
	//先入后出
	n, err := rc.Do("lpush", QueueName, uid)
	fmt.Println(n, err)
}
func startWorker() {
	//限制最大为2个协程，具体看单台处理能力而定，不要太多
	for i := 0; i < ThreadNum; i++ {
		wg.Add(1)
		go worker()
		log.Printf("start thread num : %d", i)
	}
}
func worker() {
	wg.Done()
	for {
		rc := connRedis()
		//开启后一直挂起不要退出
		uid := int64(0)
		listRes, err := redis.Strings(rc.Do("brpop", QueueName,10))
		log.Printf("uid %v done",listRes)
		if len(listRes) < 1 {
			log.Println("Queue is empty")
		}else{
			uid,_ = strconv.ParseInt(listRes[1], 0, 64)
		}
		if err != nil {
			log.Println(err)
		}
		rc.Close()
		//具体的业务逻辑
		log.Printf("uid %d done",uid)
		//如果队列是空的 则sleep 60s
		time.Sleep(SleepTime * time.Second)
	}
}
func connRedis() redis.Conn {
	rc, err := redis.Dial("tcp", RedisHost)
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return rc
	}
	fmt.Println(rc.Do("ping"))
	return rc
}

func main() {
	connRedis()
	//启动worker
	go startWorker()
	http.HandleFunc("/", IndexHandler)
	err := http.ListenAndServe(":8000", nil)
	log.Fatal(err)
}
