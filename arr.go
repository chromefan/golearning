package main

import (
	"fmt"
)

//数组3个数相加为0
func arrAdd(numArr []int) {
	l := len(numArr)
	p := l/3
	for i :=0 ;i < p; i++ {
		for j := i +1 ; j < l-2; j = j + 2{
			sumTmp := numArr[i] + numArr[j] + numArr[j+1]
			if sumTmp == 0  {
				fmt.Println(i,j,j+1)
			}
		}
	}
}

func strToInt(str string)  uint64 {
	n := uint64(0)
	max := len(str)
	for i := 0;i < max ;i++ {
		ch := str[i]-'0'
		n = 10*n+uint64(ch)
		fmt.Println(n,uint64(ch))
	}
	return n
}
func main()  {
	str := "1234"
	strToInt(str)
	fmt.Println("hello")
	arr := []int{1,3,0,4,-1,-3,0}
	arrAdd(arr)
}