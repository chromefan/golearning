package main

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

//1、snowflakes 算法得到全局唯一并自增的ID
//2、根据唯一自增的ID hash 算法为 6位

type SortUrl struct {
	Sid     int64  `bson:"sid"`
	SortUrl string `bson:"sortUrl"`
	Url     string `bson:"url"`
	UrlMd5  string `bson:"urlMd5"`
	Ctime   int64  `bson:"ctime"`
}

func main() {
	url := "https://segmentfault.com/a/1190000013831352"
	str := encode(url)
	strbytes := []byte(str)
	encoded := base64.StdEncoding.EncodeToString(strbytes)
	fmt.Println(str, encoded)
}
func encode(url string) string {
	md5Str := MD5(url)
	//1、查找这url是否存在
	if urlIsExist(md5Str) {
		return ""
	}
	//获取唯一自增ID
	id := int64(0)
	sortUrlStr := ""
	for {
		id = getId()
		//转化为62位
		sortUrlStr = DecimalToSortString(int(id), 62)
		if false == sortUrlIsExist(sortUrlStr) {
			break
		}
	}
	var sortUrl SortUrl
	sortUrl.Sid = id
	sortUrl.Url = url
	sortUrl.UrlMd5 = md5Str
	sortUrl.SortUrl = sortUrlStr
	sortUrl.Ctime = time.Now().Unix()
	insertUrl(sortUrl)
	return sortUrlStr
}
func urlIsExist(url string) bool {
	if url != "" {
		return false
	}
	return true
}
func sortUrlIsExist(sortUrl string) bool {
	if sortUrl != "" {
		return false
	}
	return true
}
func insertUrl(sortUrl SortUrl) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("sorturl").Collection("sorturl")
	fmt.Println(sortUrl)
	res, err := collection.InsertOne(context.TODO(), sortUrl)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Printf("InsertOne插入的消息ID:%v\n", res.InsertedID)
	//// _id: 默认生成一个全局唯一ID, ObjectID：12字节的二进制
	docId := res.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID:", docId.Hex())
	return true
}
func getId() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	// Generate a snowflake ID.
	id := node.Generate()
	// Print out the ID in a few different ways.
	fmt.Println(id)
	fmt.Printf("String ID: %s\n", id)
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base64 ID: %s\n", id.Base64())
	return int64(id)
}
func DecimalToSortString(num, n int) string {
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

// 生成md5
func MD5(str string) string {
	c := md5.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}
