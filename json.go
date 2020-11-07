package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"os"
	"time"
)

func main() {

	fmt.Println(CRC32("123456"))
	fmt.Println(MD5("123456"))
	fmt.Println(SHA1("123456"))
	mqid := 1558237040301332480
	fmt.Println(1558237040301332480/100000)
	mod := (mqid/10000) % 64
	fmt.Println(mod)
	return
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


// 生成md5
func MD5(str string) string {
	c := md5.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

//生成sha1
func SHA1(str string) string{
	c:=sha1.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

func CRC32(str string) uint32{
	return crc32.ChecksumIEEE([]byte(str))
}
