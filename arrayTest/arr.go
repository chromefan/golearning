package main

import "fmt"

func main() {
	var a = [3]int{1,2,3}
	defer test()
	for i, v :=range a {
		fmt.Printf("%d %d\n", i, v)
	}
	type Currency int
	const  (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD:"$",EUR:"&",GBP:"#",RMB:"¥"}
	fmt.Println(RMB,symbol[RMB])
}

func test(){
	fmt.Println("I M TEST defer")
	defer fmt.Println("Fourth")
	fmt.Println("First")
	fmt.Println("Third")
}


