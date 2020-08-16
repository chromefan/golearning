package main

import (
	"fmt"
	"time"
)
var tokens = make(chan struct{},1)
func main() {
	//存放生产的channel
	jobChan := make(chan int, 100)
	//通知是否完成所有任务
	endChan := make(chan bool)

	for i:=0; i<=10;i++  {
		go production(jobChan)
	}
	go worker(jobChan, endChan)

	select {
	case <-endChan:
		fmt.Println("消费完成……………………")
		return
	case <-time.After(time.Second * 10000):
		fmt.Println("超时………………………")
		return
	}
}

//消费
func worker(jobChan <-chan int, endChan chan bool) {
	num := 0
	for job := range jobChan {
		fmt.Println("消费:", job)
		time.Sleep(2 * time.Second)
		num++
	}
	//消费结束，通知endChan
	endChan <- true
	fmt.Println("消费总计:", num)
}

//生产
func production(jobChan chan<- int) {
	fmt.Println("生产start")
	tokens <- struct{}{}
	for i := 1; i <= 10; i++ {
		fmt.Println("生产:", i)
		jobChan <- i
	}
	fmt.Println("生产end")
	<-tokens
	//关闭channel防止消费阻塞
}
