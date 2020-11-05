package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)
	go func(ch chan int) {
		for i := 0; i < 100; i++ {
			ch <- i
			if i % 2 == 1 {
				fmt.Println("r1 : ", i)
			}
		}
	}(ch)
	go func(ch chan int) {
		for i := 0; i < 100; i++ {
			 <- ch
			if i % 2 == 0 {
				fmt.Println("r2 : ", i)
			}
		}
	}(ch)
	close(ch)
	//close(ch2)
	//time.Sleep(100 * time.Second)
}
