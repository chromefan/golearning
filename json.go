package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	nowTime := time.Now().Unix()
	timestamp := nowTime - 86400
	tm := time.Unix(timestamp, 0)
	day := tm.Format("20060102")
	var logTypeList []string
	logTypeList =append(logTypeList,"log")
	logTypeList =append(logTypeList,"log.wf")

	for _, logType := range logTypeList{
		filename := fmt.Sprintf("%s/%s.%s.%s", "log", "time", logType, day)
		fmt.Println(filename)
		err := os.Remove(filename)
		if err !=nil {
			continue
		}
	}
	err := os.Remove("test.txt")
	fmt.Println(err)
	return
	s := `[1, 2, 3, 4]`
	var a []int
	// 将字符串反解析为数组
	json.Unmarshal([]byte(s), &a)
	fmt.Println(a)  // [1 2 3 4]
}
