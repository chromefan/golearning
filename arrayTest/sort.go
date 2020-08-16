package main

import (
	"sort"
	"fmt"
)
type person struct {
	Name string
	Age int
	Id int
}

type personSlice []person

// 这三个方法必须有，相当于实现了sort.Interface
func (s personSlice) Len() int { return len(s) }
func (s personSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s personSlice) Less(i, j int) bool {
	if  s[i].Age == s[j].Age{
		return  s[i].Id > s[j].Id
	}else{
		return s[i].Age > s[j].Age
	}
}// 这里是关键，我比较了年龄这个字段

func main() {
	a := personSlice {
		{
			Name: "AAA",
			Age: 5,
			Id:1,
		},
		{
			Name: "BBB",
			Age: 5,
			Id:2,
		},
		{
			Name: "CCC",
			Age: 5,
			Id:3,
		},
		{
			Name: "DDD",
			Age: 5,
			Id:4,
		},
		{
			Name: "EEE",
			Age: 5,
			Id:5,
		},
	}
	sort.Sort(a)
	fmt.Println(a)
}