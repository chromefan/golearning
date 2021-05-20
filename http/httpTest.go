package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	ch := make(chan int)
	baseurl := "http://127.0.0.1:8089/"

	run(ch,baseurl)
	time.Sleep(3*time.Second)

}

func run(ch chan int,baseurl string) {
	num := 0
	fnum := 0
	start := time.Now().Unix()
	for i := 0; i < 10; i++ {
		url := fmt.Sprintf("%s?id=%d", baseurl, i)
		log.Println("http : ", url)
		//time.Sleep(100 * time.Microsecond)
		go func() {
			for j := 0; j < 10000; j++ {
				res, err := Get(url)
				num ++
				if err != nil {
					fnum ++
				}
				fmt.Printf("res : %s num :%d fnum: %d \n", res, num, fnum)
				return
			}
		}()
	}
	stop := time.Now().Unix()
	fmt.Println("total time :", stop-start)
	<-ch
	close(ch)
}

// 发送GET请求
// url:请求地址
// response:请求返回的内容
func Get(url string) (string, error) {
	client := http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
	}
	return result.String(), nil
}

// 发送POST请求
// url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json
// content:请求放回的内容
func Post(url string, data interface{}, contentType string) string {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest(`POST`, url, bytes.NewBuffer(jsonStr))
	req.Header.Add(`content-type`, contentType)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
