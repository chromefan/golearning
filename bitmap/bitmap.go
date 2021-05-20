package main

import (
	"github.com/spaolacci/murmur3"
	"github.com/willf/bitset"
	"fmt"
	"strconv"
)
const DEFAULT_SIZE = 2<<24
var seeds = []uint{7, 11, 13, 31, 37, 61}

type BloomFilter struct {
	set *bitset.BitSet
	funcs [6]SimpleHash
}

func NewBloomFilter() *BloomFilter {
	bf := new(BloomFilter)
	for i:=0;i< len(bf.funcs);i++{
		bf.funcs[i] = SimpleHash{DEFAULT_SIZE,seeds[i]}
	}
	bf.set = bitset.New(DEFAULT_SIZE)
	return bf
}

func (bf BloomFilter) add(value string){
	for _,f:=range(bf.funcs){
		bf.set.Set(f.hash(value))
	}
}

func (bf BloomFilter) contains(value string) bool  {
	if(value == ""){
		return false
	}
	ret := true
	for _,f:=range(bf.funcs){
		ret = ret && bf.set.Test(f.hash(value))
	}
	return ret
}


type SimpleHash struct{
	cap uint
	seed uint
}

func (s SimpleHash) hash(value string) uint{
	var result uint = 0
	for i:=0;i< len(value);i++{
		result = result*s.seed+uint(value[i])
	}
	res := (s.cap-1)&result
	fmt.Println(res)
	return (s.cap-1)&result
}
var bitmap []int
func phoneToBit(phone int)  {
	fmt.Println(strconv.IntSize)
	fmt.Println(uint(phone%(8*strconv.IntSize)))
	bitmap[phone/(8*strconv.IntSize)] |= 1 << uint(phone%(8*strconv.IntSize))
	fmt.Println(bitmap)
}

func DecimalToAny(num, n int) string {
	var tenToAny map[int]string = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: ":", 37: ";", 38: "<", 39: "=", 40: ">", 41: "?", 42: "@", 43: "[", 44: "]", 45: "^", 46: "_", 47: "{", 48: "|", 49: "}", 50: "A", 51: "B", 52: "C", 53: "D", 54: "E", 55: "F", 56: "G", 57: "H", 58: "I", 59: "J", 60: "K", 61: "L", 62: "M", 63: "N", 64: "O", 65: "P", 66: "Q", 67: "R", 68: "S", 69: "T", 70: "U", 71: "V", 72: "W", 73: "X", 74: "Y", 75: "Z"}
	new_num_str := ""
	var remainder int
	var remainder_string string
	for num != 0 {
		remainder = num % n
		if 76 > remainder && remainder > 9 {
			remainder_string = tenToAny[remainder]
		} else {
			remainder_string = strconv.Itoa(remainder)
		}
		new_num_str = remainder_string + new_num_str
		num = num / n
	}
	return new_num_str
}

func main() {
	incr := murmur3.Sum32([]byte("https://u.geekbang.org/subject/python/100038901?utm_source=wechat&utm_medium=pyq02282300&utm_term=wechatpyq02282300"))
	fmt.Println(incr)
	fmt.Println(DecimalToAny(int(incr),62))
	fmt.Println( 10 << 2)
	str :="afds"
	fmt.Println(TestHash(str))
	return
	hashStr := SDBMHash(str)
	fmt.Println(hashStr)
	hashStr = DJBHash(str)
	fmt.Println(hashStr)
	phone := 12345678
	phoneToBit(phone)
	filter := NewBloomFilter()
	fmt.Println(filter.funcs[1].seed)
	str1 := "hello,bloom filter!"
	filter.add(str1)
	return
	str2 := "A happy day"
	filter.add(str2)
	str3 := "Greate wall"
	filter.add(str3)
	fmt.Println(filter.contains(str1))
	fmt.Println(filter.contains(str2))
	fmt.Println(filter.contains(str3))
	fmt.Println(filter.contains("blockchain technology"))
}