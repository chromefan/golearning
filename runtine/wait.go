package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

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

	fmt.Println("end")
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
		str := fmt.Sprintf("%d\n",i)
		WriteWithIo("worker1.txt",str)
		time.Sleep(1*time.Second)
	}

}
func worker2(wg *sync.WaitGroup,thread int)  {
	fmt.Println("worker2 thread: ",thread)
	//str := fmt.Sprintf("%d\n",thread)
	//WriteWithIo("worker2.txt",str)
	time.Sleep(1*time.Second)
	wg.Done()
}
func WriteWithIo(filnename,content string) {
	fileObj,err := os.OpenFile(filnename,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err != nil {
		fmt.Println("Failed to open the file",err.Error())
		os.Exit(2)
	}
	if  _,err := io.WriteString(fileObj,content);err == nil {
		fmt.Println("Successful appending to the file with os.OpenFile and io.WriteString.",content)
	}
}