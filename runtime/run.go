package main

import (
	"fmt"
	"io"
	"os"
	"time"
)


func main() {
	fmt.Println("start")
	t1 := time.Now()
	workList := make(chan int,5)
	for i:=0;i<10 ;i++  {
		go add(workList)
	}
	//go dcr(workList)
	res := <- workList
	fmt.Println("end",res)
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}
func dcr(workList chan int)  {
	num := 0
	for i := 0; i < 10 ; i++ {
		str := fmt.Sprintf("%d\n",i)
		WriteWithIo("test2.txt",str)
		fmt.Println("dcr : ",str)
		time.Sleep(1)
		//ch <- i
		num +=i
	}
	workList <- num
	//defer close(workList)
}
func add(ch chan int)  {
	num :=0
	for i :=0 ; i< 10 ; i++ {
		str := fmt.Sprintf("%d\n",i)
		WriteWithIo("test1.txt",str)
		fmt.Println("add :",str)
		num += i
	}
	time.Sleep(1*time.Second)
	ch <- num
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