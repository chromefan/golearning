package main

import (
	"fmt"
	"sync"
	"time"
)

func start()  {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		println("come into goroutine1")
		for c := range ch {
			fmt.Println(c)
		}
		fmt.Println("werwer")
	}()
	fmt.Println("done ")
	defer func() {
		fmt.Println("defer done ")
	}()
	wg.Wait()
}
func main() {
	println("start main")
	start()
	time.Sleep(10*time.Second)


	/*go func() {
		println("come into goroutine2")
		var r int = 1
		for i := 1; i <= 10; i++ {
			r *= i
		}
		ch <- r
	}()

	go func() {
		println("come into goroutine3")
		ch <- 11
	}()

	for i := 0; i < 3; i++ {
		result += <-ch
	}*/
	/*//close(ch)
	println("result is:", result)
	println("end main")*/
}