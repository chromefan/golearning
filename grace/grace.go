package main

import (
	"github.com/facebookgo/grace/gracehttp"
	"net/http"
	"time"
)

func main() {
	gracehttp.Serve(
		&http.Server{Addr: ":5001", Handler: newGraceHandler()},
	)
}

func newGraceHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/sleep", func(w http.ResponseWriter, r *http.Request) {
		duration, err := time.ParseDuration(r.FormValue("duration"))
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		time.Sleep(duration)
		w.Write([]byte("Hello World"))
	})
	return mux
}
