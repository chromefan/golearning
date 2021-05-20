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

type MessageInfoMongo struct {
	ThreadId        int64       `json:"thread_id"`
	Lctime          int64       `json:"lctime"`
	LctimeMs        int64       `json:"lctime_ms"`
	Unread          string      `json:"unread"`
	Box             string      `json:"box"`
	Card            string      `json:"card"`
	Content         string      `json:"content"`
	Type            string      `json:"type"`
	Md5             string      `json:"md5"`
	Imei            interface{} `json:"imei"`
	Ctime           int64       `json:"_ctime"`
	Mtime           int64       `json:"_mtime"`
	IsDelete        int         `json:"_isdelete"`
	Key             string      `json:"_key"`
	Device          string      `json:"_device"`
	Uid             int64       `json:"uid"`
	Person          interface{} `json:"person"`
	PersonFormatted interface{} `json:"person_formatted"`
	Name            interface{} `json:"name"`
}

func main() {

	var msg MessageInfoMongo
	str := "{\"_ctime\":1612161505,\"_isdelete\":0,\"_key\":\"6e032446c222bfbc-1612161505\",\"_mtime\":1612161505,\"box\":\"receive\",\"card\":\"\",\"content\":\"【互动吧】你已成功报名《世界读书日~樊登读...》，请尽快登录互动吧App领取你的报名凭证https://hudongba.com/dw/f6gxn\",\"imei\":[\"GM22_20204409F296070CB02B00F3E64099E1\"],\"lctime\":1587007051337,\"lctime_ms\":1587007051337,\"md5\":\"3BD5C1CDABADB473CC0FBC41B66AC956\",\"name\":{\"is_encrypt\":1,\"value\":\"4e770uE4TYvsgV5DeUavivaO7ZHGZsWvWvv/vP8bODD2yfBGOlD1NNYDJE5FfQ\"},\"person\":{\"is_encrypt\":1,\"value\":\"5c3bI9iWkMeQwx1Ms7ExhEG9LAd+JcYpJ/dycwqaC1z9iPI9bu53Gj+6xrrjzA\"},\"person_formatted\":{\"is_encrypt\":1,\"value\":\"518a/yn9B/REOXMmEj76efRdY3aa5GBaYasA9e7K1TCv7Zq2LFxWe/e8wPNzMQ\"},\"thread_id\":634,\"type\":\"sms\",\"unread\":\"1\"}"

	type User struct {
		ID   int64  `json:"_id"`
		Name string `json:"name"`
	}
	var user User
	user.ID = 1
	user.Name = "test"
	userByte, err := json.Marshal(&user)
	fmt.Println(string(userByte))
	err = json.Unmarshal(userByte, &user)
	fmt.Println(err, user)
	err = json.Unmarshal([]byte(str), &msg)
	fmt.Println(err, msg)
	fmt.Println("Ctime ", msg.Ctime)
	fmt.Println("Mtime ", msg.Mtime)
	return
	fmt.Println(CRC32("123456"))
	fmt.Println(MD5("123456"))
	fmt.Println(SHA1("123456"))
	mqid := 1558237040301332480
	fmt.Println(1558237040301332480 / 100000)
	mod := (mqid / 10000) % 64
	fmt.Println(mod)
	return
	nowTime := time.Now().Unix()
	timestamp := nowTime - 86400
	tm := time.Unix(timestamp, 0)
	day := tm.Format("20060102")
	var logTypeList []string
	logTypeList = append(logTypeList, "log")
	logTypeList = append(logTypeList, "log.wf")

	for _, logType := range logTypeList {
		filename := fmt.Sprintf("%s/%s.%s.%s", "log", "time", logType, day)
		fmt.Println(filename)
		err := os.Remove(filename)
		if err != nil {
			continue
		}
	}
	err = os.Remove("test.txt")
	fmt.Println(err)
	return
	s := `[1, 2, 3, 4]`
	var a []int
	// 将字符串反解析为数组
	json.Unmarshal([]byte(s), &a)
	fmt.Println(a) // [1 2 3 4]
}

// 生成md5
func MD5(str string) string {
	c := md5.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

//生成sha1
func SHA1(str string) string {
	c := sha1.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

func CRC32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}
