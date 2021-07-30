package main

import (
	"fmt"
	"sync"
	"time"
)

var result []string
func main() {

	t :=TimeSubDay("2020-06-19","2020-01-01",)
	fmt.Println(t)
	fmt.Println("start")
	var wg sync.WaitGroup
	wg.Add(10)

	for i:=0; i<10 ; i++  {
		go worker2(&wg,i)
	}
	wg.Wait()
	fmt.Println(result)
	fmt.Println("end")
}
func worker2(wg *sync.WaitGroup,thread int)  {
	fmt.Println("worker2 thread: ",thread)
	time.Sleep(1*time.Second)
	res :=  fmt.Sprintf("i am %d",thread)
	result = append(result,res)
	wg.Done()
}
func TimeSubDay(t_str1 string, t_str2 string) int {
	layout := "2006-01-02"
	t1, _ := time.Parse(layout, t_str1)
	t2, _ := time.Parse(layout, t_str2)

	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

func worker1(thread int)  {

	fmt.Println("worker1 thread: ",thread)
	for i:=0 ; i< 10 ; i++ {
		time.Sleep(1*time.Second)
	}
	result[thread] =  fmt.Sprintf("i am %d",thread)

}
