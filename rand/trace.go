package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)


func TraceID() uint64 {
	rand.Seed(time.Now().UnixNano())
	return uint64(rand.Uint32())<<12 + uint64(rand.Uint32())
}

func IsMobileNo(str string) bool {
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	//reg := `^1(3[0-9]|5[012356789]|7[1235678]|8[0-9])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(str)
}
func main() {
	id := TraceID()
	str := fmt.Sprintf("%d",id)
	fmt.Println(id,len(str))
	str = "13812345678"
	fmt.Println("mobile",IsMobileNo(str))
}