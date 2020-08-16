package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*months := [...]string{1: "January", 12: "December"}
	for i, v := range months {
		fmt.Printf("%d:%s\n" ,i,v)
	}
	Q2 := months[4:7]
	fmt.Println(Q2)*/

	for i:=0 ; i < 10 ; i++ {
		s := rand.Intn(10)
		fmt.Printf("%d:%d\n" ,i,s)
	}
	cacheRecords := [4]int{80, 79, 78, 77}
	fmt.Println(cacheRecords[0])
	//var b []int =  //creates a slice from a[1] to a[3]
	fmt.Println(cacheRecords[2:])
	fmt.Println(cacheRecords[:2])
}
