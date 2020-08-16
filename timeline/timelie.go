package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {

	fmt.Println("start")
		//每天运行一次
	t1 := time.Now()
	for {
		var wg sync.WaitGroup
		wg.Add(2)
		for i:=0; i<2 ; i++  {
			go worker1(&wg,i)
			fmt.Println("i::",i)
		}
		wg.Wait()

		fmt.Println("end")

		elapsed := time.Since(t1)
		fmt.Println("App elapsed: ", elapsed)
		time.Sleep(1* time.Hour)
	}

}
func TimeSubDay(t_str1 string, t_str2 string) int {
	layout := "2006-01-02"
	t1, _ := time.Parse(layout, t_str1)
	t2, _ := time.Parse(layout, t_str2)

	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

func worker1(wg *sync.WaitGroup,thread int)  {

	fmt.Println("worker1 thread: ",thread)
	for i:=0 ; i< 10 ; i++ {
		str := fmt.Sprintf("%d\n",i)
		rnum := rand.Intn(10)
		fmt.Println(".........str: ",str,rnum)
		//WriteWithIo("worker1.txt",str)
		if i==1 {
			time.Sleep(10*time.Second)
		}
		time.Sleep(1*time.Second)
	}
	defer func() {
		wg.Done()
	}()

}
func worker2(wg *sync.WaitGroup,thread int)  {
	fmt.Println("worker2 thread: ",thread)
	//str := fmt.Sprintf("%d\n",thread)
	//WriteWithIo("worker2.txt",str)
	time.Sleep(1*time.Second)
	defer func() {
		wg.Done()
	}()

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