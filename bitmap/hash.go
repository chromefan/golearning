package main

import "fmt"

func SDBMHash(str string) uint64 {
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = uint64(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}
func DJBHash(str string) uint64{
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = ((hash << 5) + hash) + uint64(str[i])
	}
	return hash
}
func TestHash(str string) uint64 {
	hash := uint64(0)
	for i := 0; i < len(str); i++ {
		hash = (hash << 5) + hash + uint64(str[i])
	}
	return hash
}
func main()  {
	str :="afds234234sdfasdf"
	fmt.Println(str)
	fmt.Println("TestHash:",TestHash(str))
	hashStr := SDBMHash(str)
	fmt.Println("SDBMHash:",hashStr)
	hashStr = DJBHash(str)
	fmt.Println("DJBHash:",hashStr)
}