package main

import (
	"fmt"
)


func main() {
	ch := make(chan bool)
	 urlNum  := -1
	 go func() {
		 if urlNum > 0 {
			 fmt.Println( urlNum)
			 ch <- true
		 }
	 }()
	select {
	case <-ch:
		fmt.Println("repeat")
	default:
	 }
	fmt.Println("done")
}