package main

import "fmt"

func FindRepeatStr(str string) string {
	strList := []rune(str)
	strLen := len(str)
	hashMap := make(map[string]int)
	max := 0

	var maxStr string
	page := 0
	//fmt.Println(max)
	for i := 0; i < (strLen-1)/2; i++ {
		page = i + 2
		for j := 0; j < strLen; j += page {
			//打印2
			tmp := string(strList[j])
			for p := 1; p < page; p++ {
				if (j + p) < strLen {
					tmp += string(strList[j+p])
				}
			}
			if len(tmp) < 2 {
				continue
			}
			if _, ok := hashMap[tmp]; ok {
				hashMap[tmp]++
			} else {
				hashMap[tmp] = 1
			}
			sum := hashMap[tmp] * len(tmp)
			//fmt.Println(sum)
			if max < sum {
				//过滤小的值
				max = sum
				maxStr = tmp
			}
			fmt.Println(tmp, j)
		}
		fmt.Println("i-----------", i)
		fmt.Println("hash-----------", hashMap)
	}
	return maxStr
	for i := 0; i < strLen; {
		//打印2
		tmp := string(strList[i])
		if (i + 1) < strLen {
			tmp += string(strList[i+1])
		}
		fmt.Println(tmp, i)
		i += 2
	}
	for i := 0; i < strLen; {
		//打印2
		tmp := string(strList[i])
		if (i + 3) < strLen {
			tmp += string(strList[i+1]) + string(strList[i+2])
		}
		fmt.Println(tmp, i)
		i += 3
	}
	for i := 0; i < strLen; {
		//打印2
		tmp := string(strList[i])
		if (i + 4) < strLen {
			tmp += string(strList[i+1]) + string(strList[i+2]) + string(strList[i+3])
		}
		fmt.Println(tmp, i)
		i += 4
	}
	return str
	for i := 0; i < strLen; i++ {
		//找到重复的字串并且记录
		pre := 0

		tmp := string(strList[i])

		fmt.Println("strLen,pre, i,tmp", strLen, pre, i, tmp)
		for j := i + 1; j < (strLen-1)/2; j++ {
			//累加
			tmp += string(strList[j])
			if _, ok := hashMap[tmp]; ok {
				hashMap[tmp]++
			} else {
				hashMap[tmp] = 1
			}
			sum := hashMap[tmp] * len(tmp)
			//fmt.Println(sum)
			if max < sum {
				//过滤小的值
				max = sum
				maxStr = tmp
			}
			pre++
		}
		i += pre

	}
	fmt.Println(hashMap)
	return maxStr

}

func main() {
	str := "abcabcabab"
	//str := "abcdabc"
	resp := FindRepeatStr(str)
	fmt.Println(resp)
}
