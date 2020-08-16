package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func index(ctx *fasthttp.RequestCtx) {
	time.Sleep(200 * time.Millisecond)
	fmt.Fprintf(ctx.Response.BodyWriter(), "aaa")
}

const ConnMaxPools  = 256*1024*2

func main() {

	log.Println("fasthttp")
	handle := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			index(ctx)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}
	go func() {
		http.ListenAndServe("0.0.0.0:8099", nil)
	}()
	srv := fasthttp.Server{
		Handler: handle,
		Concurrency:ConnMaxPools,
		ReadTimeout:2000*time.Millisecond,
	}
	err := srv.ListenAndServe(":8089")
	log.Fatal(err)

}