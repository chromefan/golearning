package main

import (
	"fmt"
	"golib/pnet"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
)

type Request struct {
	Param  []byte
	Result chan []byte
}

func main() {
	pool, _ := pnet.NewPoolWithFunc(1000, func(payload interface{}) {
		request, ok := payload.(*Request)
		if !ok {
			return
		}
		reverseParam := func(s []byte) []byte {
			for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
				s[i], s[j] = s[j], s[i]
			}
			return s
		}(request.Param)
		//time.Sleep(1*time.Second)
		fmt.Printf("玩完了，下一个\n-----------------\n")
		request.Result <- reverseParam
	})
	defer pool.Release()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		param, err := ioutil.ReadAll(r.Body)
		log.Println(r.RequestURI)
		//fmt.Println( "params"+string(param))
		if err != nil {
			http.Error(w, "request error", http.StatusInternalServerError)
		}
		defer r.Body.Close()
		//log.Println(22)
		request := &Request{Param: param, Result: make(chan []byte)}

		// Throttle the requests traffic with ants pool. This process is asynchronous and
		// you can receive a result from the channel defined outside.
		if err := pool.Invoke(request); err != nil {
			http.Error(w, "throttle limit error", http.StatusInternalServerError)
		}

		fmt.Printf("正在http请求%s 还有1秒钟....\n",r.URL.Query()["id"])

		w.Write(<-request.Result)

	})

	http.ListenAndServe(":8089", nil)
}

