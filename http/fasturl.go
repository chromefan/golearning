package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
	"net/url"
	"time"
)

var proxyClient = &fasthttp.Client{
	//Addr: "10.242.110.44:8002",

	Dial: fasthttpproxy.FasthttpHTTPDialer("127.0.0.1:8888"),
}

func main() {
	urlInfo, err := url.ParseRequestURI("https://www.baidu.com/")
	if err != nil {
		fmt.Println(err)
	}
	//proxyClient.MaxIdleConnDuration = 10 * time.Millisecond
	u := urlInfo.String()
	fmt.Println(u)
	httpCode, resp, err := proxyClient.GetTimeout(nil, u, 15000 * time.Millisecond)
	fmt.Println(httpCode)
	fmt.Println(string(resp))
	fmt.Println(err)
}
