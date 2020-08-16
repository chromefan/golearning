package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://10.242.110.44:8002")
	}

	transport := &http.Transport{Proxy: proxy}

	client := &http.Client{Transport: transport}
	resp, err := client.Get("http://www.baidu.com")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
