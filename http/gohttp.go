package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(200 * time.Millisecond)
	w.Write([]byte("hi"))
}

func main() {
	http.HandleFunc("/", hiHandler)
	http.ListenAndServe(":8088", nil)
}